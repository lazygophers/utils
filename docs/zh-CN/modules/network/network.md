---
title: network
---

# network

`network` 提供**网络环境与地址相关的辅助能力**：本机 IP / 网卡探测、`net/netip` 标准库之上的 CIDR 合并与起止地址计算。

## 适合什么场景

- 启动时探测本机网络信息（IP / 网卡）。
- 把多份重叠 / 相邻的 CIDR 前缀合并为最小规范集合（防火墙规则去重、IP 段管理）。
- 从 CIDR 拿起止地址或地址总数（路由表、IP 池规划）。

## 常用入口

### IP 探测

- `IsLocalIp(ip string) bool` — 私网 / 回环 / 链路本地。
- 网卡 / fiber 辅助见同包 `interface.go` / `fiber.go`。

### CIDR 合并（基于 `net/netip`）

- `MergeCIDRs(prefixes ...netip.Prefix) []netip.Prefix` — 合并多个前缀。
- `MergeCIDRStrings(prefixes ...string) []netip.Prefix` — 字符串入口，解析失败跳过。

合并语义：
1. IPv4 / IPv6 各自独立合并，不跨族
2. 重叠 / 被包含的前缀被吸收
3. 相邻同长度兄弟合并为短一位的父前缀（迭代直到稳定）
4. 输入 invalid / zero-value 跳过；空输入返 nil

```go
network.MergeCIDRStrings(
    "10.0.0.0/25", "10.0.0.128/25",     // → 10.0.0.0/24
    "10.0.0.0/16", "10.0.5.0/24",       // → 10.0.0.0/16（包含吸收）
)
```

### CIDR 起止地址

- `CIDRStart(p netip.Prefix) netip.Addr` — 网络首址。
- `CIDREnd(p netip.Prefix) netip.Addr` — 末址（主机位全 1，v4 即广播地址）。
- `CIDRStartEnd(p) (start, end netip.Addr)` — 同时返回。
- `CIDRCount(p) *big.Int` — 地址总数（兼容 IPv6 /0 = 2^128）。

```go
p := netip.MustParsePrefix("10.0.0.0/24")
start, end := network.CIDRStartEnd(p)  // 10.0.0.0, 10.0.0.255
count := network.CIDRCount(p)          // 256
```

## 使用建议

- 网络信息高度依赖当前机器、容器、网卡和运行环境，示例结果不应当成固定行为。
- 涉及公网 / 内网、多网卡或容器网络时，要在真实部署环境验证。
- CIDR 操作完全基于 `net/netip` 标准库，无第三方依赖。
- 如果目标是 URL 规范化而不是网络探测，请看 [urlx](/modules/network/urlx)。

## 相关文档

- [urlx](/modules/network/urlx)
- [cryptox](/modules/network/cryptox)
- [系统与配置](/modules/system/)
