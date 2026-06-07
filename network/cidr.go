package network

import (
	"math/big"
	"net/netip"
	"sort"
)

// MergeCIDRs 合并多个 CIDR 前缀，返回最小规范化集合：
//   - 重叠 / 被包含的前缀被吸收
//   - 相邻同长度的兄弟前缀合并为短一位的父前缀
//   - IPv4 / IPv6 各自独立合并，不跨族
//
// 输入空或全部 invalid 时返回 nil。
// 输出按地址升序、bits 升序排列；每个 Prefix 已 Masked 规范化。
func MergeCIDRs(prefixes ...netip.Prefix) []netip.Prefix {
	v4, v6 := splitFamily(prefixes)
	merged := append(collapse(v4), collapse(v6)...)
	if len(merged) == 0 {
		return nil
	}
	return merged
}

// MergeCIDRStrings 接受字符串形式的前缀（如 "10.0.0.0/24"），解析失败的条目跳过，
// 其余走 MergeCIDRs 合并。
func MergeCIDRStrings(prefixes ...string) []netip.Prefix {
	parsed := make([]netip.Prefix, 0, len(prefixes))
	for _, s := range prefixes {
		p, err := netip.ParsePrefix(s)
		if err != nil {
			continue
		}
		parsed = append(parsed, p)
	}
	return MergeCIDRs(parsed...)
}

// splitFamily 按 IPv4 / IPv6 分组，丢弃 invalid。
func splitFamily(prefixes []netip.Prefix) (v4, v6 []netip.Prefix) {
	for _, p := range prefixes {
		if !p.IsValid() {
			continue
		}
		canon := p.Masked()
		if canon.Addr().Is4() {
			v4 = append(v4, canon)
		} else {
			v6 = append(v6, canon)
		}
	}
	return v4, v6
}

// collapse 对单族（已 Masked）做"包含吸收 + 兄弟合并"循环直到稳定。
func collapse(list []netip.Prefix) []netip.Prefix {
	if len(list) == 0 {
		return nil
	}
	sort.Slice(list, func(i, j int) bool {
		c := list[i].Addr().Compare(list[j].Addr())
		if c != 0 {
			return c < 0
		}
		return list[i].Bits() < list[j].Bits()
	})

	list = absorbContained(list)

	for {
		next, changed := mergeSiblings(list)
		if !changed {
			return next
		}
		list = next
	}
}

// absorbContained 假定已排序：前一个若覆盖后一个，则吸收后者。
func absorbContained(list []netip.Prefix) []netip.Prefix {
	out := list[:0]
	var prev netip.Prefix
	for _, p := range list {
		if prev.IsValid() && prev.Bits() <= p.Bits() && prev.Contains(p.Addr()) {
			continue
		}
		out = append(out, p)
		prev = p
	}
	return out
}

// mergeSiblings 扫一遍把相邻同长度兄弟合并为父前缀；返回新切片与是否发生合并。
func mergeSiblings(list []netip.Prefix) ([]netip.Prefix, bool) {
	if len(list) < 2 {
		return list, false
	}
	out := make([]netip.Prefix, 0, len(list))
	changed := false
	i := 0
	for i < len(list) {
		if i+1 < len(list) && areSiblings(list[i], list[i+1]) {
			parent := netip.PrefixFrom(list[i].Addr(), list[i].Bits()-1).Masked()
			out = append(out, parent)
			i += 2
			changed = true
			continue
		}
		out = append(out, list[i])
		i++
	}
	return out, changed
}

// areSiblings 判定两同长度前缀是否同父（去掉最末一位后相等）。
func areSiblings(a, b netip.Prefix) bool {
	if a.Bits() != b.Bits() || a.Bits() == 0 {
		return false
	}
	if a.Addr().Is4() != b.Addr().Is4() {
		return false
	}
	parentBits := a.Bits() - 1
	return netip.PrefixFrom(a.Addr(), parentBits).Masked() == netip.PrefixFrom(b.Addr(), parentBits).Masked()
}

// CIDRStart 返回 CIDR 前缀的起始地址（网络地址）；invalid 输入返回零值 Addr。
func CIDRStart(p netip.Prefix) netip.Addr {
	if !p.IsValid() {
		return netip.Addr{}
	}
	return p.Masked().Addr()
}

// CIDREnd 返回 CIDR 前缀的结束地址（主机位全 1）；invalid 输入返回零值 Addr。
// IPv4 即广播地址；IPv6 即区间末址。
func CIDREnd(p netip.Prefix) netip.Addr {
	if !p.IsValid() {
		return netip.Addr{}
	}
	canon := p.Masked()
	addr := canon.Addr()
	bits := canon.Bits()

	if addr.Is4() {
		b := addr.As4()
		setHostBits(b[:], bits)
		return netip.AddrFrom4(b)
	}
	b := addr.As16()
	setHostBits(b[:], bits)
	return netip.AddrFrom16(b)
}

// CIDRStartEnd 同时返回 CIDR 前缀的起止地址。
func CIDRStartEnd(p netip.Prefix) (start, end netip.Addr) {
	return CIDRStart(p), CIDREnd(p)
}

// CIDRCount 返回前缀覆盖的地址总数；用 *big.Int 兼容 IPv6（/0 时 2^128）。
// invalid 输入返回 nil。
func CIDRCount(p netip.Prefix) *big.Int {
	if !p.IsValid() {
		return nil
	}
	hostBits := p.Addr().BitLen() - p.Bits()
	return new(big.Int).Lsh(big.NewInt(1), uint(hostBits))
}

// setHostBits 把 b 中第 prefix 位起的所有位置 1。
func setHostBits(b []byte, prefix int) {
	total := len(b) * 8
	if prefix >= total {
		return
	}
	// 前缀所在字节的低位部分
	byteIdx := prefix / 8
	bitInByte := prefix % 8
	if bitInByte != 0 {
		b[byteIdx] |= 0xFF >> uint(bitInByte)
		byteIdx++
	}
	// 剩余字节整字节填 0xFF
	for i := byteIdx; i < len(b); i++ {
		b[i] = 0xFF
	}
}
