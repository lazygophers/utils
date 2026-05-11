package validator

import (
	"testing"
	"unicode"
)

// 测试数据
var (
	validStrongPasswords = []string{
		"Abc123!@",           // 最小长度，包含所有类型
		"Password123!",       // 常见强密码
		"SecurePass#2024",    // 包含数字
		"MyP@ssw0rd",         // 复杂密码
		"Test@1234",          // 简单但符合
		"ADMIN@123",          // 全大写+数字+特殊
		"student#123",        // 全小写+数字+特殊
		"User2024$Pass",      // 混合
		"1A2b3C4d!",          // 交替字符
		"P@ssw0rd123456",     // 较长密码
	}
	invalidStrongPasswords = []string{
		"",              // 空
		"short1A!",      // 太短
		"nocaps123!",    // 无大写
		"NOLOWER123!",   // 无小写
		"NoNumber!!",    // 无数字
		"NoSpecial123",  // 无特殊字符
		"onlylower",     // 只小写
			"ONLYUPPER",    // 只大写
		"12345678",      // 只数字
		"!@#$%^&*",      // 只特殊字符
		"Abc123",        // 长度够但只2种类型
		"ABCDEFGH",      // 只1种类型
	}
	weakPasswords = []string{
		"Password1",     // 无特殊字符（2种类型）
		"password!",     // 无数字（2种类型）
		"PASSWORD1",     // 无小写和特殊（2种类型）
		"12345678a",     // 无大写和特殊（2种类型）
	}
)

// ============== 方案1：当前实现（基线） ==============
func validateStrongPasswordOriginal(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasNumber = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 方案2：使用字节循环替代 rune ==============
func validateStrongPasswordByteLoop(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for i := 0; i < len(password); i++ {
		c := password[i]
		switch {
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		case c >= 'a' && c <= 'z':
			hasLower = true
		case c >= '0' && c <= '9':
			hasNumber = true
		case (c >= '!' && c <= '/') || (c >= ':' && c <= '@') || (c >= '[' && c <= '`') || (c >= '{' && c <= '~'):
			hasSpecial = true
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 方案3：使用查找表优化 ASCII 范围检查 ==============
func validateStrongPasswordLookupTable(password string) bool {
	if len(password) < 8 {
		return false
	}

	// 预定义字符类型查找表
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for i := 0; i < len(password); i++ {
		c := password[i]
		switch {
		case (c >= 'A' && c <= 'Z'):
			hasUpper = true
		case (c >= 'a' && c <= 'z'):
			hasLower = true
		case (c >= '0' && c <= '9'):
			hasNumber = true
		default:
			// 其他 ASCII 字符视为特殊字符
			if c >= 32 && c <= 126 {
				hasSpecial = true
			}
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 方案4：快速失败优化 ==============
func validateStrongPasswordFastFail(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   uint8
		hasLower   uint8
		hasNumber  uint8
		hasSpecial uint8
	)

	for i := 0; i < len(password); i++ {
		c := password[i]
		switch {
		case c >= 'A' && c <= 'Z':
			hasUpper = 1
		case c >= 'a' && c <= 'z':
			hasLower = 1
		case c >= '0' && c <= '9':
			hasNumber = 1
		default:
			if c >= 32 && c <= 126 {
				hasSpecial = 1
			}
		}

		// 快速失败：已经找到3种类型且遍历了足够长度
		if hasUpper+hasLower+hasNumber+hasSpecial >= 3 {
			return true
		}
	}

	return hasUpper+hasLower+hasNumber+hasSpecial >= 3
}

// ============== 方案5：位掩码优化 ==============
func validateStrongPasswordBitMask(password string) bool {
	if len(password) < 8 {
		return false
	}

	const (
		upperMask   = 1 << 0
		lowerMask   = 1 << 1
		numberMask  = 1 << 2
		specialMask = 1 << 3
	)

	var mask uint8

	for i := 0; i < len(password); i++ {
		c := password[i]
		switch {
		case c >= 'A' && c <= 'Z':
			mask |= upperMask
		case c >= 'a' && c <= 'z':
			mask |= lowerMask
		case c >= '0' && c <= '9':
			mask |= numberMask
		default:
			if c >= 32 && c <= 126 {
				mask |= specialMask
			}
		}
	}

	// 检查是否至少有3种类型
	return (mask&upperMask + mask&lowerMask + mask&numberMask + mask&specialMask) >= 3
}

// ============== 方案6：预计算特殊字符范围 ==============
func validateStrongPasswordPrecompute(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for i := 0; i < len(password); i++ {
		c := password[i]
		switch {
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		case c >= 'a' && c <= 'z':
			hasLower = true
		case c >= '0' && c <= '9':
			hasNumber = true
		case c >= 33 && c <= 47, c >= 58 && c <= 64, c >= 91 && c <= 96, c >= 123 && c <= 126:
			hasSpecial = true
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 方案7：分支消除优化 ==============
func validateStrongPasswordBranchless(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   uint8
		hasLower   uint8
		hasNumber  uint8
		hasSpecial uint8
	)

	for i := 0; i < len(password); i++ {
		c := password[i]
		upper := uint8(0)
		if c >= 'A' && c <= 'Z' {
			upper = 1
		}
		lower := uint8(0)
		if c >= 'a' && c <= 'z' {
			lower = 1
		}
		digit := uint8(0)
		if c >= '0' && c <= '9' {
			digit = 1
		}
		special := uint8(0)
		if c >= 33 && c <= 126 {
			if c < '0' || c > '9' {
				if c < 'A' || c > 'Z' {
					if c < 'a' || c > 'z' {
						special = 1
					}
				}
			}
		}

		hasUpper += upper
		hasLower += lower
		hasNumber += digit
		hasSpecial += special
	}

	count := uint8(0)
	if hasUpper > 0 {
		count++
	}
	if hasLower > 0 {
		count++
	}
	if hasNumber > 0 {
		count++
	}
	if hasSpecial > 0 {
		count++
	}

	return count >= 3
}

// ============== 方案8：ASCII 表查找 ==============
func validateStrongPasswordASCIITable(password string) bool {
	if len(password) < 8 {
		return false
	}

	// 字符类型表：0=其他, 1=大写, 2=小写, 3=数字, 4=特殊
	var charType [128]uint8
	for c := 'A'; c <= 'Z'; c++ {
		charType[c] = 1
	}
	for c := 'a'; c <= 'z'; c++ {
		charType[c] = 2
	}
	for c := '0'; c <= '9'; c++ {
		charType[c] = 3
	}
	// 特殊字符
	for c := '!'; c <= '/'; c++ {
		charType[c] = 4
	}
	for c := ':'; c <= '@'; c++ {
		charType[c] = 4
	}
	for c := '['; c <= '`'; c++ {
		charType[c] = 4
	}
	for c := '{'; c <= '~'; c++ {
		charType[c] = 4
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for i := 0; i < len(password); i++ {
		c := password[i]
		if c < 128 {
			switch charType[c] {
			case 1:
				hasUpper = true
			case 2:
				hasLower = true
			case 3:
				hasNumber = true
			case 4:
				hasSpecial = true
			}
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 方案9：混合 Unicode + ASCII 快速路径 ==============
func validateStrongPasswordHybrid(password string) bool {
	if len(password) < 8 {
		return false
	}

	// 先检查是否为纯 ASCII
	isASCII := true
	for i := 0; i < len(password); i++ {
		if password[i] > 127 {
			isASCII = false
			break
		}
	}

	if isASCII {
		// 快速 ASCII 路径
		var (
			hasUpper   bool
			hasLower   bool
			hasNumber  bool
			hasSpecial bool
		)

		for i := 0; i < len(password); i++ {
			c := password[i]
			switch {
			case c >= 'A' && c <= 'Z':
				hasUpper = true
			case c >= 'a' && c <= 'z':
				hasLower = true
			case c >= '0' && c <= '9':
				hasNumber = true
			default:
				if c >= 32 && c <= 126 {
					hasSpecial = true
				}
			}
		}

		count := 0
		if hasUpper {
			count++
		}
		if hasLower {
			count++
		}
		if hasNumber {
			count++
		}
		if hasSpecial {
			count++
		}

		return count >= 3
	}

	// Unicode 路径（原始实现）
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 方案10：计数器累加优化 ==============
func validateStrongPasswordCounter(password string) bool {
	if len(password) < 8 {
		return false
	}

	count := 0

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for i := 0; i < len(password); i++ {
		c := password[i]
		switch {
		case c >= 'A' && c <= 'Z':
			if !hasUpper {
				hasUpper = true
				count++
			}
		case c >= 'a' && c <= 'z':
			if !hasLower {
				hasLower = true
				count++
			}
		case c >= '0' && c <= '9':
			if !hasNumber {
				hasNumber = true
				count++
			}
		default:
			if c >= 32 && c <= 126 && !hasSpecial {
				hasSpecial = true
				count++
			}
		}

		// 快速失败
		if count >= 3 {
			return true
		}
	}

	return count >= 3
}

// ============== 方案11：SIMD 风格批量处理 ==============
func validateStrongPasswordSIMDStyle(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	i := 0
	// 批量处理 4 个字符
	for i+4 <= len(password) {
		c1, c2, c3, c4 := password[i], password[i+1], password[i+2], password[i+3]

		// 字符 1
		switch {
		case c1 >= 'A' && c1 <= 'Z':
			hasUpper = true
		case c1 >= 'a' && c1 <= 'z':
			hasLower = true
		case c1 >= '0' && c1 <= '9':
			hasNumber = true
		default:
			if c1 >= 32 && c1 <= 126 {
				hasSpecial = true
			}
		}

		// 字符 2
		switch {
		case c2 >= 'A' && c2 <= 'Z':
			hasUpper = true
		case c2 >= 'a' && c2 <= 'z':
			hasLower = true
		case c2 >= '0' && c2 <= '9':
			hasNumber = true
		default:
			if c2 >= 32 && c2 <= 126 {
				hasSpecial = true
			}
		}

		// 字符 3
		switch {
		case c3 >= 'A' && c3 <= 'Z':
			hasUpper = true
		case c3 >= 'a' && c3 <= 'z':
			hasLower = true
		case c3 >= '0' && c3 <= '9':
			hasNumber = true
		default:
			if c3 >= 32 && c3 <= 126 {
				hasSpecial = true
			}
		}

		// 字符 4
		switch {
		case c4 >= 'A' && c4 <= 'Z':
			hasUpper = true
		case c4 >= 'a' && c4 <= 'z':
			hasLower = true
		case c4 >= '0' && c4 <= '9':
			hasNumber = true
		default:
			if c4 >= 32 && c4 <= 126 {
				hasSpecial = true
			}
		}

		i += 4
	}

	// 处理剩余字符
	for i < len(password) {
		c := password[i]
		switch {
		case c >= 'A' && c <= 'Z':
			hasUpper = true
		case c >= 'a' && c <= 'z':
			hasLower = true
		case c >= '0' && c <= '9':
			hasNumber = true
		default:
			if c >= 32 && c <= 126 {
				hasSpecial = true
			}
		}
		i++
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 方案12：内联优化 ==============
func validateStrongPasswordInlined(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for i := 0; i < len(password); i++ {
		c := password[i]
		if c >= 'A' && c <= 'Z' {
			hasUpper = true
		} else if c >= 'a' && c <= 'z' {
			hasLower = true
		} else if c >= '0' && c <= '9' {
			hasNumber = true
		} else if c >= 32 && c <= 126 {
			hasSpecial = true
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasNumber {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}

// ============== 基准测试 ==============

// BenchmarkStrongPassword_Original 当前实现（基线）
func BenchmarkStrongPassword_Original(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordOriginal(pwd)
			}
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, pwd := range invalidStrongPasswords {
				_ = validateStrongPasswordOriginal(pwd)
			}
		}
	})

	b.Run("Weak", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, pwd := range weakPasswords {
				_ = validateStrongPasswordOriginal(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_ByteLoop 字节循环
func BenchmarkStrongPassword_ByteLoop(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordByteLoop(pwd)
			}
		}
	})

	b.Run("Invalid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range invalidStrongPasswords {
				_ = validateStrongPasswordByteLoop(pwd)
			}
		}
	})

	b.Run("Weak", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range weakPasswords {
				_ = validateStrongPasswordByteLoop(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_LookupTable 查找表优化
func BenchmarkStrongPassword_LookupTable(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordLookupTable(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_FastFail 快速失败
func BenchmarkStrongPassword_FastFail(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordFastFail(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_BitMask 位掩码
func BenchmarkStrongPassword_BitMask(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordBitMask(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_Precompute 预计算
func BenchmarkStrongPassword_Precompute(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordPrecompute(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_Branchless 分支消除
func BenchmarkStrongPassword_Branchless(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordBranchless(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_ASCIITable ASCII 表
func BenchmarkStrongPassword_ASCIITable(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordASCIITable(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_Hybrid 混合
func BenchmarkStrongPassword_Hybrid(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordHybrid(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_Counter 计数器
func BenchmarkStrongPassword_Counter(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordCounter(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_SIMDStyle SIMD 风格
func BenchmarkStrongPassword_SIMDStyle(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordSIMDStyle(pwd)
			}
		}
	})
}

// BenchmarkStrongPassword_Inlined 内联优化
func BenchmarkStrongPassword_Inlined(b *testing.B) {
	b.Run("Valid", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			for _, pwd := range validStrongPasswords {
				_ = validateStrongPasswordInlined(pwd)
			}
		}
	})
}
