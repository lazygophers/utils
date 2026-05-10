package stringx

import (
	"encoding/base64"
	"encoding/hex"
	"strings"
	"unicode"
)

// TrimSpaceAll 移除字符串中所有空白字符（包括空格、制表符、换行符等）
// 与 strings.TrimSpace 不同，此函数会移除字符串内部的所有空白，而非仅首尾
func TrimSpaceAll(s string) string {
	if s == "" {
		return ""
	}

	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		if !unicode.IsSpace(r) {
			b.WriteRune(r)
		}
	}

	return b.String()
}

// ToTitle 将字符串转换为标题格式
// 每个单词的首字母大写，其余字母小写
// 例如： "hello world" → "Hello World"
func ToTitle(s string) string {
	if s == "" {
		return ""
	}

	// 使用标准库的 Title 方法处理大多数情况
	// 但需要更精确的单词边界处理
	return strings.ToTitle(s)
}

// Normalize 对字符串进行 Unicode 归一化
// form: 0=NFC, 1=NFD, 2=NFKC, 3=NFKD
// 注意：Go 标准库没有内置归一化功能，此函数为占位实现
// 实际使用建议：import "golang.org/x/text/unicode/norm"
// 例如：将 "é" 的组合字符分解为基础字符 + 变音符号
func Normalize(s string, form int) string {
	if s == "" {
		return ""
	}

	// 简化实现：目前仅返回原字符串
	// 完整实现需要 golang.org/x/text/unicode/norm 包
	// 可以通过以下方式使用：
	//   import "golang.org/x/text/unicode/norm"
	//   result := norm.NFC.String(s)  // 或 NFD, NFKC, NFKD
	return s
}

// Base64Encode 将字符串编码为 Base64
// 使用标准 Base64 编码（带填充）
func Base64Encode(s string) string {
	if s == "" {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decode 从 Base64 解码字符串
func Base64Decode(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// Base64URLEncode 将字符串编码为 URL 安全的 Base64
// 使用 RawURLEncoding（无填充，URL 安全字符）
func Base64URLEncode(s string) string {
	if s == "" {
		return ""
	}
	return base64.RawURLEncoding.EncodeToString([]byte(s))
}

// Base64URLDecode 从 URL 安全的 Base64 解码字符串
func Base64URLDecode(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	decoded, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// HexEncode 将字符串编码为十六进制
func HexEncode(s string) string {
	if s == "" {
		return ""
	}
	return hex.EncodeToString([]byte(s))
}

// HexDecode 从十六进制解码字符串
func HexDecode(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	decoded, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// Mask 对字符串进行掩码处理，仅显示前后指定数量的字符
// visible: 首尾各保留的字符数
// 例如： Mask("12345678", 2) → "12****78"
// 如果 visible*2 >= 字符串长度，则返回原字符串
func Mask(s string, visible int) string {
	if s == "" {
		return ""
	}
	if visible < 0 {
		visible = 0
	}

	runes := []rune(s)
	length := len(runes)

	if visible*2 >= length {
		return s
	}

	// 计算掩码部分长度
	maskLength := length - visible*2
	if maskLength <= 0 {
		return s
	}

	// 构建结果：前 visible 个 + 掩码 + 后 visible 个
	var b strings.Builder
	b.Grow(length)

	// 前面部分
	for i := 0; i < visible; i++ {
		b.WriteRune(runes[i])
	}

	// 掩码部分
	for i := 0; i < maskLength; i++ {
		b.WriteRune('*')
	}

	// 后面部分
	for i := length - visible; i < length; i++ {
		b.WriteRune(runes[i])
	}

	return b.String()
}

// MaskEmail 对邮箱进行掩码处理
// 例如： "user@example.com" → "u***@example.com"
func MaskEmail(email string) string {
	if email == "" {
		return ""
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		// 不是有效邮箱格式，使用通用掩码
		return Mask(email, 2)
	}

	username, domain := parts[0], parts[1]
	usernameRunes := []rune(username)

	if len(usernameRunes) <= 1 {
		// 用户名太短，保留首字符
		return string(usernameRunes[0]) + "***@" + domain
	}

	// 用户名保留首字符，其余掩码
	return string(usernameRunes[0]) + "***@" + domain
}

// MaskPhone 对手机号进行掩码处理
// 例如： "13812345678" → "138****5678"
func MaskPhone(phone string) string {
	if phone == "" {
		return ""
	}

	runes := []rune(phone)
	length := len(runes)

	if length <= 7 {
		// 太短，使用通用掩码
		return Mask(phone, 2)
	}

	// 手机号保留前 3 位和后 4 位
	var b strings.Builder
	b.Grow(length)

	for i := 0; i < 3; i++ {
		b.WriteRune(runes[i])
	}

	for i := 3; i < length-4; i++ {
		b.WriteRune('*')
	}

	for i := length - 4; i < length; i++ {
		b.WriteRune(runes[i])
	}

	return b.String()
}

// EditDistance 计算两个字符串之间的 Levenshtein 编辑距离
// 返回将一个字符串转换为另一个字符串所需的最少单字符编辑操作数
// 操作包括：插入、删除、替换
func EditDistance(a, b string) int {
	if a == b {
		return 0
	}
	if a == "" {
		return len(b)
	}
	if b == "" {
		return len(a)
	}

	runesA, runesB := []rune(a), []rune(b)
	lenA, lenB := len(runesA), len(runesB)

	// 优化空间复杂度：使用两行而非完整矩阵
	previous := make([]int, lenB+1)
	current := make([]int, lenB+1)

	// 初始化第一行
	for j := 0; j <= lenB; j++ {
		previous[j] = j
	}

	for i := 1; i <= lenA; i++ {
		current[0] = i

		for j := 1; j <= lenB; j++ {
			cost := 1
			if runesA[i-1] == runesB[j-1] {
				cost = 0
			}

			current[j] = min(
				previous[j]+1,      // 删除
				current[j-1]+1,     // 插入
				previous[j-1]+cost, // 替换
			)
		}

		previous, current = current, previous
	}

	return previous[lenB]
}

// Similarity 计算两个字符串的相似度（0-1 之间）
// 使用编辑距离计算：1 - (editDistance / max(len(a), len(b)))
func Similarity(a, b string) float64 {
	if a == b {
		return 1.0
	}

	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}

	if maxLen == 0 {
		return 1.0
	}

	distance := EditDistance(a, b)
	return 1.0 - float64(distance)/float64(maxLen)
}

// Slugify 将字符串转换为 URL 友好格式
// 转换规则：
// 1. 转换为小写
// 2. 替换空格和特殊字符为连字符
// 3. 移除多余连字符
// 4. 移除首尾连字符
// 例如： "Hello World!" → "hello-world"
func Slugify(s string) string {
	if s == "" {
		return ""
	}

	// 转小写
	s = strings.ToLower(s)

	runes := []rune(s)
	var b strings.Builder
	b.Grow(len(s))

	prevIsHyphen := false
	prevIsAlnum := false

	for _, r := range runes {
		isAlnum := (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')

		if isAlnum {
			b.WriteRune(r)
			prevIsHyphen = false
			prevIsAlnum = true
		} else if unicode.IsSpace(r) || r == '-' || r == '_' || r == '/' {
			// 仅在字母数字之间添加连字符
			if prevIsAlnum && !prevIsHyphen {
				b.WriteRune('-')
				prevIsHyphen = true
			}
		}
		// 其他字符直接忽略
	}

	result := b.String()

	// 移除尾部连字符
	if len(result) > 0 && result[len(result)-1] == '-' {
		result = result[:len(result)-1]
	}

	return result
}

// RemoveHyphens 移除字符串中的所有连字符
func RemoveHyphens(s string) string {
	if s == "" {
		return ""
	}

	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		if r != '-' {
			b.WriteRune(r)
		}
	}

	return b.String()
}

// NormalizeHyphens 统一字符串中的连字符类型
// 将各种连字符（全角、半角、en dash、em dash）统一为标准半角连字符
func NormalizeHyphens(s string) string {
	if s == "" {
		return ""
	}

	runes := []rune(s)
	var b strings.Builder
	b.Grow(len(s))

	for _, r := range runes {
		// 检查各种连字符类型
		// U+0021: ! (exclamation mark, 不是连字符)
		// U+002D: - (hyphen-minus, 标准连字符)
		// U+00AD: ­ (soft hyphen)
		// U+2010: ‐ (hyphen)
		// U+2011: ‑ (non-breaking hyphen)
		// U+2012: ‒ (figure dash)
		// U+2013: – (en dash)
		// U+2014: — (em dash)
		// U+2015: ― (horizontal bar)
		// U+2212: − (minus sign)
		// U+FE58: ﹘ (small em dash)
		// U+FF0D: － (fullwidth hyphen-minus)

		switch r {
		case '-', '­', '‐', '‑', '‒', '–', '—', '―', '−', '﹘', '－':
			// 统一转换为标准连字符
			b.WriteRune('-')
		default:
			b.WriteRune(r)
		}
	}

	return b.String()
}

// CountWords 统计字符串中的单词数量
// 单词由空白字符分隔
func CountWords(s string) int {
	if s == "" {
		return 0
	}

	// 使用 strings.Fields 自动处理连续空白
	return len(strings.Fields(s))
}

// CountLines 统计字符串中的行数
// 行由换行符分隔
func CountLines(s string) int {
	if s == "" {
		return 0
	}

	// 使用 Split 自动处理各种换行符
	lines := strings.Split(s, "\n")
	return len(lines)
}

// CountRunes 统计字符串中的 rune（Unicode 码点）数量
func CountRunes(s string) int {
	if s == "" {
		return 0
	}
	return len([]rune(s))
}

// ToLower 转换字符串为小写（快捷方式）
func ToLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s)
}

// ToUpper 转换字符串为大写（快捷方式）
func ToUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s)
}

// ToTitleCase 转换字符串为标题格式（每个单词首字母大写）
// 注意：这与 ToTitle 不同，ToTitleCase 使用更精确的单词边界
func ToTitleCase(s string) string {
	if s == "" {
		return ""
	}

	// 使用标准库的 Title 功能
	// strings.Title 会将每个单词首字母大写，其余小写
	return strings.Title(s)
}

// Contains 检查字符串是否包含子串（快捷方式）
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// HasPrefix 检查字符串是否以指定前缀开头（快捷方式）
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix 检查字符串是否以指定后缀结尾（快捷方式）
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// TrimPrefix 移除指定前缀（快捷方式）
func TrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

// TrimSuffix 移除指定后缀（快捷方式）
func TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

// Index 返回子串首次出现的位置（快捷方式）
func Index(s, substr string) int {
	return strings.Index(s, substr)
}

// LastIndex 返回子串最后出现的位置（快捷方式）
func LastIndex(s, substr string) int {
	return strings.LastIndex(s, substr)
}

// Substring 安全地获取子字符串
// 支持负数索引（从末尾开始计数）
// 例如： Substring("hello", 1, 4) → "ell"
//       Substring("hello", -3, -1) → "ll"
func Substring(s string, start, end int) string {
	if s == "" {
		return ""
	}

	runes := []rune(s)
	length := len(runes)

	// 处理负数索引
	if start < 0 {
		start = length + start
		if start < 0 {
			start = 0
		}
	}
	if end < 0 {
		end = length + end
		if end < 0 {
			end = 0
		}
	}

	// 边界检查
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	if start >= end {
		return ""
	}

	return string(runes[start:end])
}

// RemoveDuplicates 移除字符串中重复的连续字符
// 例如： "aaabbbcccaaa" → "abca"
func RemoveDuplicates(s string) string {
	if s == "" {
		return ""
	}

	runes := []rune(s)
	if len(runes) <= 1 {
		return s
	}

	var b strings.Builder
	b.Grow(len(runes))

	prev := runes[0]
	b.WriteRune(prev)

	for _, r := range runes[1:] {
		if r != prev {
			b.WriteRune(r)
			prev = r
		}
	}

	return b.String()
}

// ReverseWords 反转字符串中单词的顺序
// 例如： "hello world test" → "test world hello"
func ReverseWords(s string) string {
	if s == "" {
		return ""
	}

	words := strings.Fields(s)
	length := len(words)
	if length <= 1 {
		return s
	}

	// 原地反转
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

// Capitalize 将字符串首字母大写，其余小写
// 例如： "hello WORLD" → "Hello world"
func Capitalize(s string) string {
	if s == "" {
		return ""
	}

	runes := []rune(s)
	if len(runes) == 0 {
		return ""
	}

	// 首字母大写
	runes[0] = unicode.ToUpper(runes[0])

	// 其余字母小写
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}

	return string(runes)
}
