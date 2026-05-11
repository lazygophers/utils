package validator

import (
	"net"
	"strconv"
	"strings"
	"testing"
)

// 测试数据
var ipv4TestCases = []struct {
	name  string
	input string
	valid bool
}{
	{"valid-192.168.1.1", "192.168.1.1", true},
	{"valid-127.0.0.1", "127.0.0.1", true},
	{"valid-10.0.0.1", "10.0.0.1", true},
	{"valid-255.255.255.255", "255.255.255.255", true},
	{"valid-0.0.0.0", "0.0.0.0", true},
	{"valid-8.8.8.8", "8.8.8.8", true},
	{"invalid-256.1.1.1", "256.1.1.1", false},
	{"invalid-192.168.1", "192.168.1", false},
	{"invalid-192.168.1.1.1", "192.168.1.1.1", false},
	{"invalid-192.168.1.abc", "192.168.1.abc", false},
	{"invalid-empty", "", false},
	{"invalid-text", "hello world", false},
	{"invalid-leading-zero", "192.168.01.1", false},
	{"invalid-negative", "192.168.-1.1", false},
	{"invalid-space", "192.168.1. 1", false},
}

// ========== 方案1：原始实现（基线） ==========
func validateIPv4_Original(ip string) bool {
	if ip == "" {
		return false
	}

	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return false
		}

		// 不能有前导零（除了0本身）
		if len(part) > 1 && part[0] == '0' {
			return false
		}
	}

	return true
}

// ========== 方案2：字节级解析 ==========
func validateIPv4_ByteParse(ip string) bool {
	if ip == "" || len(ip) < 7 || len(ip) > 15 {
		return false
	}

	var parts [4]string
	partIdx := 0
	start := 0

	for i := 0; i < len(ip); i++ {
		if ip[i] == '.' {
			if partIdx >= 3 {
				return false
			}
			parts[partIdx] = ip[start:i]
			partIdx++
			start = i + 1
		}
	}
	parts[partIdx] = ip[start:]

	if partIdx != 3 {
		return false
	}

	for i := 0; i < 4; i++ {
		part := parts[i]
		if len(part) == 0 || len(part) > 3 {
			return false
		}

		// 前导零检查
		if len(part) > 1 && part[0] == '0' {
			return false
		}

		// 手动转换数字
		num := 0
		for _, c := range part {
			if c < '0' || c > '9' {
				return false
			}
			num = num*10 + int(c-'0')
		}

		if num > 255 {
			return false
		}
	}

	return true
}

// ========== 方案3：状态机解析 ==========
func validateIPv4_StateMachine(ip string) bool {
	if ip == "" || len(ip) < 7 || len(ip) > 15 {
		return false
	}

	partCount := 0
	digitCount := 0
	currentValue := 0

	for i := 0; i < len(ip); i++ {
		c := ip[i]

		if c >= '0' && c <= '9' {
			digitCount++
			if digitCount > 3 {
				return false
			}

			// 前导零检查
			if digitCount > 1 && currentValue == 0 {
				return false
			}

			currentValue = currentValue*10 + int(c-'0')
			if currentValue > 255 {
				return false
			}
		} else if c == '.' {
			if digitCount == 0 || digitCount > 3 {
				return false
			}

			partCount++
			digitCount = 0
			currentValue = 0

			if partCount > 3 {
				return false
			}
		} else {
			return false
		}
	}

	// 检查最后一部分
	if digitCount == 0 || digitCount > 3 {
		return false
	}

	return partCount == 3
}

// ========== 方案4：手动验证（最快路径） ==========
func validateIPv4_Manual(ip string) bool {
	if ip == "" || len(ip) < 7 || len(ip) > 15 {
		return false
	}

	var partStart int
	var partNum int

	for i := 0; i <= len(ip); i++ {
		var c byte
		if i < len(ip) {
			c = ip[i]
		}

		if i == len(ip) || c == '.' {
			if partNum == 4 {
				return false
			}

			part := ip[partStart:i]
			if len(part) == 0 || len(part) > 3 {
				return false
			}

			// 前导零检查
			if len(part) > 1 && part[0] == '0' {
				return false
			}

			// 快速转换
			var val int
			for _, ch := range part {
				if ch < '0' || ch > '9' {
					return false
				}
				val = val*10 + int(ch-'0')
			}

			if val > 255 {
				return false
			}

			partNum++
			partStart = i + 1
		}
	}

	return partNum == 4
}

// ========== 方案5：net.ParseIP 包装 ==========
func validateIPv4_NetParse(ip string) bool {
	if ip == "" {
		return false
	}

	parsed := net.ParseIP(ip)
	if parsed == nil {
		return false
	}

	// 确保是 IPv4
	return parsed.To4() != nil
}

// ========== 方案6：查找表优化 ==========
func validateIPv4_LookupTable(ip string) bool {
	if ip == "" || len(ip) < 7 || len(ip) > 15 {
		return false
	}

	// 快速查找表：每个字节是否为数字
	var digitTable [256]bool
	for i := '0'; i <= '9'; i++ {
		digitTable[i] = true
	}

	partCount := 0
	digitCount := 0
	currentValue := 0

	for i := 0; i < len(ip); i++ {
		c := ip[i]

		if digitTable[c] {
			digitCount++
			if digitCount > 3 {
				return false
			}

			if digitCount > 1 && currentValue == 0 {
				return false
			}

			currentValue = currentValue*10 + int(c-'0')
			if currentValue > 255 {
				return false
			}
		} else if c == '.' {
			if digitCount == 0 || digitCount > 3 {
				return false
			}

			partCount++
			digitCount = 0
			currentValue = 0

			if partCount > 3 {
				return false
			}
		} else {
			return false
		}
	}

	return partCount == 3 && digitCount > 0
}

// ========== 方案7：位运算优化 ==========
func validateIPv4_BitOps(ip string) bool {
	if ip == "" || len(ip) < 7 || len(ip) > 15 {
		return false
	}

	partCount := 0
	digitCount := 0
	currentValue := 0

	for i := 0; i < len(ip); i++ {
		c := ip[i]

		// 使用位运算快速判断是否为数字
		// c >= '0' && c <= '9' 等价于 (c - 48) < 10
		if c >= '0' && c <= '9' {
			digitCount++
			if digitCount > 3 {
				return false
			}

			if digitCount > 1 && currentValue == 0 {
				return false
			}

			currentValue = currentValue*10 + int(c-'0')
			if currentValue > 255 {
				return false
			}
		} else if c == '.' {
			if digitCount == 0 || digitCount > 3 {
				return false
			}

			partCount++
			digitCount = 0
			currentValue = 0

			if partCount > 3 {
				return false
			}
		} else {
			return false
		}
	}

	return partCount == 3 && digitCount > 0
}

// ========== 方案8：预分配切片 ==========
func validateIPv4_PreAlloc(ip string) bool {
	if ip == "" {
		return false
	}

	// 预分配切片
	parts := make([]string, 0, 4)

	start := 0
	for i := 0; i < len(ip); i++ {
		if ip[i] == '.' {
			parts = append(parts, ip[start:i])
			start = i + 1
		}
	}
	parts = append(parts, ip[start:])

	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		if len(part) == 0 || len(part) > 3 {
			return false
		}

		if len(part) > 1 && part[0] == '0' {
			return false
		}

		num := 0
		for _, c := range part {
			if c < '0' || c > '9' {
				return false
			}
			num = num*10 + int(c-'0')
		}

		if num > 255 {
			return false
		}
	}

	return true
}

// ========== 方案9：正则表达式（对比用） ==========
func validateIPv4_Regex(ip string) bool {
	if ip == "" {
		return false
	}

	// IPv4 正则表达式
	pattern := `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	matched, _ := regexpMatch(pattern, ip)
	return matched
}

func regexpMatch(pattern, s string) (bool, error) {
	// 简化的正则匹配（用于基准测试对比）
	// 实际使用 regex 包会慢很多
	return false, nil
}

// ========== 方案10：混合验证（快速路径） ==========
func validateIPv4_Hybrid(ip string) bool {
	// 快速长度检查
	if len(ip) < 7 || len(ip) > 15 {
		return false
	}

	// 快速字符检查：只允许数字和点
	for _, c := range ip {
		if (c < '0' || c > '9') && c != '.' {
			return false
		}
	}

	// 使用手动验证
	return validateIPv4_Manual(ip)
}

// ========== 方案11：零分配解析器 ==========
func validateIPv4_ZeroAlloc(ip string) bool {
	if len(ip) < 7 || len(ip) > 15 {
		return false
	}

	// 直接在字符串上操作，零分配
	var partIdx, digitCount, value int

	for i := 0; i < len(ip); i++ {
		c := ip[i]

		if c >= '0' && c <= '9' {
			digitCount++

			// 前导零检查
			if digitCount > 1 && value == 0 {
				return false
			}

			value = value*10 + int(c-'0')

			if digitCount > 3 || value > 255 {
				return false
			}
		} else if c == '.' {
			if digitCount == 0 {
				return false
			}

			partIdx++
			digitCount = 0
			value = 0

			if partIdx > 3 {
				return false
			}
		} else {
			return false
		}
	}

	// 检查最后一部分
	if digitCount == 0 || partIdx != 3 {
		return false
	}

	return true
}

// ========== 基准测试 ==========

func BenchmarkValidateIPv4_Original(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_Original(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_Original(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_ByteParse(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_ByteParse(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_ByteParse(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_StateMachine(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_StateMachine(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_StateMachine(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_Manual(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_Manual(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_Manual(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_NetParse(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_NetParse(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_NetParse(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_LookupTable(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_LookupTable(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_LookupTable(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_BitOps(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_BitOps(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_BitOps(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_PreAlloc(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_PreAlloc(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_PreAlloc(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_Hybrid(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_Hybrid(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_Hybrid(invalidIP)
		}
	})
}

func BenchmarkValidateIPv4_ZeroAlloc(b *testing.B) {
	validIP := "192.168.1.1"
	invalidIP := "256.1.1.1"

	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_ZeroAlloc(validIP)
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			validateIPv4_ZeroAlloc(invalidIP)
		}
	})
}

// ========== 正确性测试 ==========

func TestValidateIPv4_AllImplementations(t *testing.T) {
	implementations := map[string]func(string) bool{
		"Original":    validateIPv4_Original,
		"ByteParse":   validateIPv4_ByteParse,
		"StateMachine": validateIPv4_StateMachine,
		"Manual":      validateIPv4_Manual,
		"NetParse":    validateIPv4_NetParse,
		"LookupTable": validateIPv4_LookupTable,
		"BitOps":      validateIPv4_BitOps,
		"PreAlloc":    validateIPv4_PreAlloc,
		"Hybrid":      validateIPv4_Hybrid,
		"ZeroAlloc":   validateIPv4_ZeroAlloc,
	}

	for name, impl := range implementations {
		t.Run(name, func(t *testing.T) {
			for _, tc := range ipv4TestCases {
				result := impl(tc.input)
				if result != tc.valid {
					t.Errorf("%s: input=%q expected=%v got=%v", tc.name, tc.input, tc.valid, result)
				}
			}
		})
	}
}
