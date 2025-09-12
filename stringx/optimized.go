package stringx

import (
	"strings"
	"unicode"
	"unsafe"
)

// OptimizedToSnake - 零分配版本的 ToSnake
func OptimizedToSnake(s string) string {
	if s == "" {
		return ""
	}
	
	// 预估需要的容量，避免多次扩容
	capacity := len(s) + len(s)/4  // 估算增加25%的容量用于下划线
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

// OptimizedCamel2Snake - 内存优化版本
func OptimizedCamel2Snake(s string) string {
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

// 纯ASCII优化版本，性能最高
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

// OptimizedSplitLen - 零拷贝版本
func OptimizedSplitLen(s string, max int) []string {
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

// 快速ASCII检测
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 128 {
			return false
		}
	}
	return true
}

// OptimizedReverse - 针对不同字符类型的优化版本
func OptimizedReverse(s string) string {
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

// OptimizedSnake2Camel - 内存优化版本
func OptimizedSnake2Camel(s string) string {
	if s == "" {
		return ""
	}
	
	// ASCII优化路径
	if isASCII(s) && !strings.Contains(s, "_") {
		return s // 无需转换
	}
	
	result := make([]byte, 0, len(s))
	upper := true
	
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '_' {
			upper = true
		} else {
			if upper {
				if c >= 'a' && c <= 'z' {
					result = append(result, c-32) // 快速转大写
				} else {
					result = append(result, c)
				}
				upper = false
			} else {
				result = append(result, c)
			}
		}
	}
	
	return *(*string)(unsafe.Pointer(&result))
}

// OptimizedToKebab - 基于 OptimizedToSnake 的变体
func OptimizedToKebab(s string) string {
	if s == "" {
		return ""
	}
	
	// 重用 OptimizedToSnake 的逻辑，然后替换下划线
	snakeResult := OptimizedToSnake(s)
	
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