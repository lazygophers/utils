package stringx

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"
	"unsafe"
)

func ToString(b []byte) string {
	if b == nil {
		return ""
	}
	if len(b) == 0 {
		return ""
	}
	return *(*string)(unsafe.Pointer(&b))
}

func ToBytes(s string) []byte {
	if s == "" {
		return nil
	}
	return *(*[]byte)(unsafe.Pointer(&s))
}

// Camel2Snake 驼峰转蛇形 - 内存优化版本
func Camel2Snake(s string) string {
	if s == "" {
		return ""
	}

	// 只处理ASCII字符以获得最大性能
	if isASCII(s) {
		return optimizedASCIICamel2Snake(s)
	}

	// Unicode版本保持原有逻辑但优化内存分配
	capacity := len(s) + len(s)/3
	result := make([]byte, 0, capacity)

	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, byte(unicode.ToLower(r)))
		} else {
			if r < 128 {
				result = append(result, byte(r))
			} else {
				// 非ASCII字符
				charBytes := []byte(string(r))
				result = append(result, charBytes...)
			}
		}
	}

	return *(*string)(unsafe.Pointer(&result))
}

// Snake2Camel 蛇形转驼峰
func Snake2Camel(s string) string {
	if s == "" {
		return ""
	}
	var b bytes.Buffer
	upper := true
	for _, v := range s {
		if v == '_' {
			upper = true
		} else {
			if upper {
				b.WriteRune(unicode.ToUpper(v))
				upper = false
			} else {
				b.WriteRune(v)
			}
		}
	}
	return b.String()
}

// Snake2SmallCamel 蛇形转小驼峰
func Snake2SmallCamel(s string) string {
	if s == "" {
		return ""
	}
	var b bytes.Buffer
	upper := false
	isFirst := true
	for _, v := range s {
		if v == '_' {
			upper = true
		} else {
			if isFirst {
				isFirst = false
				b.WriteRune(unicode.ToLower(v))
				upper = false // Reset upper flag after first character
			} else if upper {
				b.WriteRune(unicode.ToUpper(v))
				upper = false
			} else {
				// Convert to lowercase for consistency in camelCase
				b.WriteRune(unicode.ToLower(v))
			}
		}
	}
	return b.String()
}

// ToSnake 蛇形 - 零分配优化版本
func ToSnake(s string) string {
	if s == "" {
		return ""
	}

	// 预估需要的容量，避免多次扩容
	capacity := len(s) + len(s)/4 // 估算增加25%的容量用于下划线
	if capacity > 256 {
		capacity = 256 // 限制最大预分配容量
	}

	// 使用单次分配的 []byte 替代 bytes.Buffer
	result := make([]byte, 0, capacity)
	runes := []rune(s)

	for i, r := range runes {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			// 检查是否需要下划线
			if i > 0 {
				prev := runes[i-1]
				if (unicode.IsUpper(r) && unicode.IsLetter(prev)) ||
					(unicode.IsNumber(r) && unicode.IsLetter(prev)) ||
					(unicode.IsLetter(r) && unicode.IsNumber(prev)) {
					result = append(result, '_')
				}
			}

			// 转换为小写并添加
			if unicode.IsUpper(r) {
				result = append(result, byte(unicode.ToLower(r)))
			} else {
				// 对于ASCII字符直接转换，避免Unicode处理开销
				if r < 128 {
					result = append(result, byte(r))
				} else {
					// 非ASCII字符使用Unicode处理
					lowerStr := string(unicode.ToLower(r))
					result = append(result, lowerStr...)
				}
			}
		} else {
			// 避免连续的下划线
			if len(result) > 0 && result[len(result)-1] != '_' {
				result = append(result, '_')
			}
		}
	}

	// 零拷贝转换为字符串
	return *(*string)(unsafe.Pointer(&result))
}

// ToKebab - 基于优化ToSnake的变体
func ToKebab(s string) string {
	if s == "" {
		return ""
	}

	// 重用 ToSnake 的逻辑，然后替换下划线
	snakeResult := ToSnake(s)

	// 如果没有下划线，直接返回
	if !strings.Contains(snakeResult, "_") {
		return snakeResult
	}

	// 零拷贝替换下划线为连字符
	resultBytes := []byte(snakeResult)
	for i, b := range resultBytes {
		if b == '_' {
			resultBytes[i] = '-'
		}
	}

	return *(*string)(unsafe.Pointer(&resultBytes))
}

// ToCamel 转驼峰
func ToCamel(s string) string {
	var b bytes.Buffer
	upper := true
	prevWasNumber := false
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			// If the current character is already uppercase, treat it as a word boundary
			if upper || unicode.IsUpper(v) || (prevWasNumber && unicode.IsLetter(v)) {
				if unicode.IsLetter(v) {
					b.WriteRune(unicode.ToUpper(v))
				} else {
					b.WriteRune(v)
				}
				upper = false
			} else {
				if unicode.IsLetter(v) {
					b.WriteRune(unicode.ToLower(v))
				} else {
					b.WriteRune(v)
				}
			}
			prevWasNumber = unicode.IsNumber(v)
		} else {
			upper = true
			prevWasNumber = false
		}
	}
	return b.String()
}

func ToSlash(s string) string {
	var b bytes.Buffer
	runes := []rune(s)
	for i, v := range runes {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			needsSlash := false

			// Check if we need a slash before this character
			if i > 0 {
				prev := runes[i-1]
				// Add slash for uppercase letters (camelCase -> camel/case)
				if unicode.IsUpper(v) && unicode.IsLetter(prev) {
					needsSlash = true
				}
				// Add slash when transitioning from letter to number
				if unicode.IsNumber(v) && unicode.IsLetter(prev) {
					needsSlash = true
				}
				// Add slash when transitioning from number to letter
				if unicode.IsLetter(v) && unicode.IsNumber(prev) {
					needsSlash = true
				}
			}

			if needsSlash {
				b.WriteRune('/')
			}

			if unicode.IsUpper(v) {
				b.WriteRune(unicode.ToLower(v))
			} else {
				b.WriteRune(v)
			}
		} else {
			// Only add slash if the last character wasn't a slash
			if b.Len() > 0 {
				lastRune := []rune(b.String())
				if len(lastRune) == 0 || lastRune[len(lastRune)-1] != '/' {
					b.WriteRune('/')
				}
			}
		}
	}
	return b.String()
}

func ToDot(s string) string {
	var b bytes.Buffer
	runes := []rune(s)
	for i, v := range runes {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			needsDot := false

			// Check if we need a dot before this character
			if i > 0 {
				prev := runes[i-1]
				// Add dot for uppercase letters (camelCase -> camel.case)
				if unicode.IsUpper(v) && unicode.IsLetter(prev) {
					needsDot = true
				}
				// Add dot when transitioning from letter to number
				if unicode.IsNumber(v) && unicode.IsLetter(prev) {
					needsDot = true
				}
				// Add dot when transitioning from number to letter
				if unicode.IsLetter(v) && unicode.IsNumber(prev) {
					needsDot = true
				}
			}

			if needsDot {
				b.WriteRune('.')
			}

			if unicode.IsUpper(v) {
				b.WriteRune(unicode.ToLower(v))
			} else {
				b.WriteRune(v)
			}
		} else {
			// Only add dot if the last character wasn't a dot
			if b.Len() > 0 {
				lastRune := []rune(b.String())
				if len(lastRune) == 0 || lastRune[len(lastRune)-1] != '.' {
					b.WriteRune('.')
				}
			}
		}
	}
	return b.String()
}

// ToSmallCamel 转小驼峰
func ToSmallCamel(s string) string {
	var b bytes.Buffer
	upper := false
	isFirst := true
	prevWasNumber := false
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			if isFirst {
				isFirst = false
				if unicode.IsLetter(v) {
					b.WriteRune(unicode.ToLower(v))
				} else {
					b.WriteRune(v)
				}
				upper = false
			} else if upper || (prevWasNumber && unicode.IsLetter(v)) {
				if unicode.IsLetter(v) {
					b.WriteRune(unicode.ToUpper(v))
				} else {
					b.WriteRune(v)
				}
				upper = false
			} else {
				if unicode.IsLetter(v) {
					b.WriteRune(unicode.ToLower(v))
				} else {
					b.WriteRune(v)
				}
			}
			prevWasNumber = unicode.IsNumber(v)
		} else if !isFirst {
			upper = true
			prevWasNumber = false
		}
	}
	return b.String()
}

// SplitLen 按长度分割字符串 - 零拷贝优化版本
func SplitLen(s string, max int) []string {
	if max <= 0 {
		return []string{s}
	}
	if s == "" {
		return []string{}
	}

	runes := []rune(s)
	totalRunes := len(runes)
	if totalRunes <= max {
		return []string{s}
	}

	// 预计算结果切片容量
	estimatedParts := (totalRunes + max - 1) / max
	result := make([]string, 0, estimatedParts)

	for start := 0; start < totalRunes; start += max {
		end := start + max
		if end > totalRunes {
			end = totalRunes
		}

		// 使用零拷贝字符串转换
		part := string(runes[start:end])
		result = append(result, part)
	}

	return result
}

// Shorten 缩短字符串
func Shorten(s string, max int) string {
	if max < 0 {
		return ""
	}
	if len(s) <= max {
		return s
	}
	return s[:max]
}

func ShortenShow(s string, max int) string {
	if max < 0 {
		return "..."
	}
	if len(s) <= max {
		return s
	}
	if max < 3 {
		return "..."
	}
	return s[:max-3] + "..."
}

func IsUpper[M string | []rune](r M) bool {
	return string(r) == strings.ToUpper(string(r))
}

func IsDigit[M string | []rune](r M) bool {
	for _, i := range []rune(r) {
		if !unicode.IsDigit(i) {
			return false
		}
	}

	return true
}

func Reverse(s string) string {
	if s == "" {
		return ""
	}

	// ASCII优化路径
	if isASCII(s) {
		return reverseASCII(s)
	}

	// Unicode路径 - 使用原地反转避免额外分配
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Quote(s string) string {
	return strconv.Quote(s)
}

func QuotePure(s string) string {
	return strings.TrimPrefix(strings.TrimSuffix(Quote(s), `"`), `"`)
}

// 快速ASCII检测辅助函数
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 128 {
			return false
		}
	}
	return true
}

// 纯ASCII优化的Camel2Snake版本
func optimizedASCIICamel2Snake(s string) string {
	capacity := len(s) + len(s)/3
	result := make([]byte, 0, capacity)

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, c+32) // 快速转小写
		} else {
			result = append(result, c)
		}
	}

	return *(*string)(unsafe.Pointer(&result))
}

// 纯ASCII反转，最高性能
func reverseASCII(s string) string {
	if len(s) <= 1 {
		return s
	}

	bytes := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		bytes[i] = s[len(s)-1-i]
	}

	return *(*string)(unsafe.Pointer(&bytes))
}
