package validator

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

// 测试用例
var (
	validEmails   = []string{"user@example.com", "test.user@domain.co.uk", "admin123@mail-server.org"}
	invalidEmails = []string{
		"",                    // 空
		"invalid",             // 无@
		"@example.com",        // 无本地部分
		"user@",               // 无域名
		"user@@",              // 双@
		"user@@example.com",   // 双@
		"user..name@example.com", // 连续点
		".user@example.com",   // 开头点
		"user.@example.com",   // 结尾点
		"user@.com",           // 域名开头点
		"user@domain.",        // 域名结尾点
		"user@domain..com",    // 域名连续点
		"user@domain",         // 无TLD
		"用户@example.com",    // 非ASCII
	}
)

// 方案1: 当前正则（基线）
func validateEmailRegex(email string) bool {
	if email == "" {
		return false
	}
	return emailRegex.MatchString(email)
}

// 方案2: 手动@分割验证
func validateEmailSplit(email string) bool {
	if email == "" {
		return false
	}

	// 必须包含且仅包含一个@
	atIndex := strings.Index(email, "@")
	if atIndex == -1 || strings.LastIndex(email, "@") != atIndex {
		return false
	}

	localPart := email[:atIndex]
	domain := email[atIndex+1:]

	// 本地部分非空
	if len(localPart) == 0 {
		return false
	}

	// 域名非空且必须包含.
	if len(domain) == 0 || !strings.Contains(domain, ".") {
		return false
	}

	// 检查本地部分字符
	for _, c := range localPart {
		if !isValidLocalChar(c) {
			return false
		}
	}

	// 检查域名部分字符
	for _, c := range domain {
		if !isValidDomainChar(c) {
			return false
		}
	}

	return true
}

func isValidLocalChar(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '.' || c == '_' || c == '%' ||
		c == '+' || c == '-'
}

func isValidDomainChar(c rune) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '.' || c == '-'
}

// 方案3: 字符遍历+状态机
func validateEmailStateMachine(email string) bool {
	if email == "" {
		return false
	}

	state := 0 // 0: 本地部分, 1: 域名, 2: TLD
	hasAt := false
	hasDotInDomain := false

	for i, c := range email {
		switch state {
		case 0: // 本地部分
			if c == '@' {
				if i == 0 { // @不能在开头
					return false
				}
				state = 1
				hasAt = true
			} else if !isValidLocalChar(c) {
				return false
			}
		case 1: // 域名
			if c == '.' {
				hasDotInDomain = true
			} else if !isValidDomainChar(c) {
				return false
			}
		}
	}

	return hasAt && hasDotInDomain
}

// 方案4: bytes.LastIndexByte快速查找@
func validateEmailLastIndex(email string) bool {
	if email == "" {
		return false
	}

	// 快速查找最后一个@
	atIndex := strings.LastIndexByte(email, '@')
	if atIndex == -1 {
		return false
	}

	// 检查是否只有一个@
	if strings.IndexByte(email, '@') != atIndex {
		return false
	}

	localPart := email[:atIndex]
	domain := email[atIndex+1:]

	// 基本长度检查
	if len(localPart) == 0 || len(domain) < 4 { // domain至少 x.xx
		return false
	}

	// 域名必须包含点
	dotIndex := strings.LastIndexByte(domain, '.')
	if dotIndex == -1 || dotIndex == 0 || dotIndex == len(domain)-1 {
		return false
	}

	// 快速字符类检查
	for i := 0; i < len(localPart); i++ {
		c := localPart[i]
		isValid := (c >= 'a' && c <= 'z') ||
			(c >= 'A' && c <= 'Z') ||
			(c >= '0' && c <= '9') ||
			c == '.' || c == '_' || c == '%' ||
			c == '+' || c == '-'
		if !isValid {
			return false
		}
	}

	for i := 0; i < len(domain); i++ {
		c := domain[i]
		isValid := (c >= 'a' && c <= 'z') ||
			(c >= 'A' && c <= 'Z') ||
			(c >= '0' && c <= '9') ||
			c == '.' || c == '-'
		if !isValid {
			return false
		}
	}

	return true
}

// 方案5: 前缀字符类快速检查
func validateEmailFastPath(email string) bool {
	if email == "" {
		return false
	}

	// 快速失败：长度检查
	if len(email) < 6 { // a@b.co 最短
		return false
	}

	// 快速失败：首字符必须是有效字符
	first := email[0]
	if !isValidLocalByte(first) {
		return false
	}

	// 查找@
	atIndex := strings.IndexByte(email, '@')
	if atIndex == -1 || atIndex == 0 {
		return false
	}

	// 检查只有一个@
	if strings.Count(email, "@") != 1 {
		return false
	}

	domain := email[atIndex+1:]
	if len(domain) < 4 {
		return false
	}

	// 域名必须包含点且TLD至少2字符
	lastDot := strings.LastIndexByte(domain, '.')
	if lastDot == -1 || lastDot == 0 || lastDot >= len(domain)-2 {
		return false
	}

	return true
}

func isValidLocalByte(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '.' || c == '_' || c == '%' ||
		c == '+' || c == '-'
}

// 方案6: 分段验证（本地部分+域名）
func validateEmailSegmented(email string) bool {
	if email == "" {
		return false
	}

	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		return false
	}

	// 检查只有一个@
	if strings.Index(email[atIndex+1:], "@") != -1 {
		return false
	}

	localPart := email[:atIndex]
	domain := email[atIndex+1:]

	// 验证本地部分
	if !validateLocalPart(localPart) {
		return false
	}

	// 验证域名部分
	if !validateDomainPart(domain) {
		return false
	}

	return true
}

func validateLocalPart(local string) bool {
	if len(local) == 0 {
		return false
	}

	// 检查不以点开头或结尾
	if local[0] == '.' || local[len(local)-1] == '.' {
		return false
	}

	// 检查不包含连续点
	if strings.Contains(local, "..") {
		return false
	}

	// 检查所有字符有效
	for _, c := range local {
		if !isValidLocalChar(c) {
			return false
		}
	}

	return true
}

func validateDomainPart(domain string) bool {
	if len(domain) == 0 {
		return false
	}

	// 必须包含点
	lastDot := strings.LastIndex(domain, ".")
	if lastDot == -1 {
		return false
	}

	// TLD至少2字符
	if len(domain)-lastDot-1 < 2 {
		return false
	}

	// 检查不以点开头或结尾
	if domain[0] == '.' || domain[len(domain)-1] == '.' {
		return false
	}

	// 检查不包含连续点
	if strings.Contains(domain, "..") {
		return false
	}

	// 检查所有字符有效
	for _, c := range domain {
		if !isValidDomainChar(c) {
			return false
		}
	}

	return true
}

// 方案7: RFC 5322简化版
func validateEmailRFCSimplified(email string) bool {
	if email == "" {
		return false
	}

	// 基本结构检查
	atIndex := strings.Index(email, "@")
	if atIndex <= 0 || atIndex == len(email)-1 {
		return false
	}

	localPart := email[:atIndex]
	domain := email[atIndex+1:]

	// 本地部分：1-64字符，允许字母数字和._%+-
	if len(localPart) == 0 || len(localPart) > 64 {
		return false
	}

	// 域名：基本检查
	if len(domain) < 4 || len(domain) > 255 {
		return false
	}

	// 域名必须包含点且TLD至少2字符
	dotParts := strings.Split(domain, ".")
	if len(dotParts) < 2 {
		return false
	}

	tld := dotParts[len(dotParts)-1]
	if len(tld) < 2 {
		return false
	}

	// 检查本地部分字符
	for _, c := range localPart {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) &&
			c != '.' && c != '_' && c != '%' && c != '+' && c != '-' {
			return false
		}
	}

	// 检查域名字符
	for _, c := range domain {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) &&
			c != '.' && c != '-' {
			return false
		}
	}

	return true
}

// 方案8: 避免正则，用strings.Contains
func validateEmailStrings(email string) bool {
	if email == "" {
		return false
	}

	// 必须包含@
	if !strings.Contains(email, "@") {
		return false
	}

	// 只能有一个@
	if strings.Count(email, "@") != 1 {
		return false
	}

	atIndex := strings.Index(email, "@")
	localPart := email[:atIndex]
	domain := email[atIndex+1:]

	// 本地部分和域名都非空
	if len(localPart) == 0 || len(domain) == 0 {
		return false
	}

	// 域名必须包含点
	if !strings.Contains(domain, ".") {
		return false
	}

	// 域名点后至少2字符
	lastDot := strings.LastIndex(domain, ".")
	if len(domain)-lastDot-1 < 2 {
		return false
	}

	// 简单字符检查
	for _, c := range email {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) &&
			c != '@' && c != '.' && c != '_' &&
			c != '%' && c != '+' && c != '-' {
			return false
		}
	}

	return true
}

// 方案9: 组合优化（@位置+长度+字符检查）
func validateEmailCombined(email string) bool {
	if email == "" {
		return false
	}

	// 快速长度检查
	l := len(email)
	if l < 6 || l > 254 { // RFC最大长度
		return false
	}

	// 查找@并检查位置
	atIndex := strings.IndexByte(email, '@')
	if atIndex == -1 || atIndex == 0 || atIndex > l-4 {
		return false
	}

	// 确保只有一个@
	if strings.IndexByte(email[atIndex+1:], '@') != -1 {
		return false
	}

	// 提取本地部分和域名
	localPart := email[:atIndex]
	domain := email[atIndex+1:]

	// 本地部分长度检查（RFC 64字符）
	if len(localPart) == 0 || len(localPart) > 64 {
		return false
	}

	// 域名基本检查
	if len(domain) < 4 {
		return false
	}

	// 域名必须包含点且TLD≥2字符
	dotIndex := strings.LastIndexByte(domain, '.')
	if dotIndex == -1 || dotIndex == 0 || dotIndex > len(domain)-3 {
		return false
	}

	// 字符范围检查（ASCII优化）
	for i := 0; i < len(localPart); i++ {
		c := localPart[i]
		if !isASCIILetterDigit(c) && c != '.' && c != '_' &&
			c != '%' && c != '+' && c != '-' {
			return false
		}
	}

	for i := 0; i < len(domain); i++ {
		c := domain[i]
		if !isASCIILetterDigit(c) && c != '.' && c != '-' {
			return false
		}
	}

	return true
}

func isASCIILetterDigit(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

// 方案10: 纯ASCII范围检查
func validateEmailASCII(email string) bool {
	if email == "" {
		return false
	}

	l := len(email)
	if l < 6 || l > 254 {
		return false
	}

	atIndex := -1
	dotCount := 0

	for i := 0; i < l; i++ {
		c := email[i]

		if c == '@' {
			if atIndex != -1 { // 多个@
				return false
			}
			atIndex = i
		} else if c == '.' {
			dotCount++
		}
	}

	if atIndex == -1 || atIndex == 0 || atIndex == l-1 {
		return false
	}

	if dotCount < 1 { // 域名必须有至少一个点
		return false
	}

	// 检查本地部分
	for i := 0; i < atIndex; i++ {
		if !isASCIILocalChar(email[i]) {
			return false
		}
	}

	// 检查域名部分
	for i := atIndex + 1; i < l; i++ {
		if !isASCIIDomainChar(email[i]) {
			return false
		}
	}

	return true
}

func isASCIILocalChar(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '.' || c == '_' || c == '%' ||
		c == '+' || c == '-'
}

func isASCIIDomainChar(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '.' || c == '-'
}

// 方案11: 标准库验证
func validateEmailStdLib(email string) bool {
	if email == "" {
		return false
	}
	return emailRegex.MatchString(email)
}

// 方案12: 极简验证（仅基本格式）
func validateEmailMinimal(email string) bool {
	if email == "" {
		return false
	}

	at := strings.Index(email, "@")
	if at == -1 || at == 0 || at == len(email)-1 {
		return false
	}

	domain := email[at+1:]
	if strings.LastIndexByte(domain, '.') == -1 {
		return false
	}

	return true
}

// ========== 基准测试 ==========

func BenchmarkValidateEmailRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailRegex(email)
		}
		for _, email := range invalidEmails {
			validateEmailRegex(email)
		}
	}
}

func BenchmarkValidateEmailSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailSplit(email)
		}
		for _, email := range invalidEmails {
			validateEmailSplit(email)
		}
	}
}

func BenchmarkValidateEmailStateMachine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailStateMachine(email)
		}
		for _, email := range invalidEmails {
			validateEmailStateMachine(email)
		}
	}
}

func BenchmarkValidateEmailLastIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailLastIndex(email)
		}
		for _, email := range invalidEmails {
			validateEmailLastIndex(email)
		}
	}
}

func BenchmarkValidateEmailFastPath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailFastPath(email)
		}
		for _, email := range invalidEmails {
			validateEmailFastPath(email)
		}
	}
}

func BenchmarkValidateEmailSegmented(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailSegmented(email)
		}
		for _, email := range invalidEmails {
			validateEmailSegmented(email)
		}
	}
}

func BenchmarkValidateEmailRFCSimplified(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailRFCSimplified(email)
		}
		for _, email := range invalidEmails {
			validateEmailRFCSimplified(email)
		}
	}
}

func BenchmarkValidateEmailStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailStrings(email)
		}
		for _, email := range invalidEmails {
			validateEmailStrings(email)
		}
	}
}

func BenchmarkValidateEmailCombined(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailCombined(email)
		}
		for _, email := range invalidEmails {
			validateEmailCombined(email)
		}
	}
}

func BenchmarkValidateEmailASCII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailASCII(email)
		}
		for _, email := range invalidEmails {
			validateEmailASCII(email)
		}
	}
}

func BenchmarkValidateEmailStdLib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailStdLib(email)
		}
		for _, email := range invalidEmails {
			validateEmailStdLib(email)
		}
	}
}

func BenchmarkValidateEmailMinimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, email := range validEmails {
			validateEmailMinimal(email)
		}
		for _, email := range invalidEmails {
			validateEmailMinimal(email)
		}
	}
}

// ========== 功能正确性测试 ==========

func TestValidateEmailCorrectness(t *testing.T) {
	validators := map[string]func(string) bool{
		"方案1-正则":        validateEmailRegex,
		"方案2-分割":        validateEmailSplit,
		"方案3-状态机":       validateEmailStateMachine,
		"方案4-LastIndex":  validateEmailLastIndex,
		"方案5-快速路径":      validateEmailFastPath,
		"方案6-分段":        validateEmailSegmented,
		"方案7-RFC简化":    validateEmailRFCSimplified,
		"方案8-字符串":        validateEmailStrings,
		"方案9-组合优化":      validateEmailCombined,
		"方案10-ASCII":    validateEmailASCII,
		"方案12-极简":       validateEmailMinimal,
	}

	for name, validator := range validators {
		t.Run(name, func(t *testing.T) {
			// 测试有效邮箱
			for _, email := range validEmails {
				if !validator(email) {
					t.Errorf("有效邮箱被拒绝: %s", email)
				}
			}

			// 测试无效邮箱
			for _, email := range invalidEmails {
				if validator(email) {
					t.Errorf("无效邮箱被接受: %s", email)
				}
			}
		})
	}
}

// ========== 额外性能测试：分离有效/无效 ==========

func BenchmarkValidEmailsOnly(b *testing.B) {
	validators := map[string]func(string) bool{
		"方案1-正则": validateEmailRegex,
		"方案4-LastIndex": validateEmailLastIndex,
		"方案9-组合优化":   validateEmailCombined,
		"方案10-ASCII":  validateEmailASCII,
	}

	for name, fn := range validators {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, email := range validEmails {
					fn(email)
				}
			}
		})
	}
}

func BenchmarkInvalidEmailsOnly(b *testing.B) {
	validators := map[string]func(string) bool{
		"方案1-正则": validateEmailRegex,
		"方案4-LastIndex": validateEmailLastIndex,
		"方案9-组合优化":   validateEmailCombined,
		"方案10-ASCII":  validateEmailASCII,
	}

	for name, fn := range validators {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, email := range invalidEmails {
					fn(email)
				}
			}
		})
	}
}

// 生成详细报告
func TestGenerateEmailReport(t *testing.T) {
	results := []struct {
		name  string
		valid int
		total int
	}{
		{"方案1-正则", 0, 0},
		{"方案2-分割", 0, 0},
		{"方案3-状态机", 0, 0},
		{"方案4-LastIndex", 0, 0},
		{"方案5-快速路径", 0, 0},
		{"方案6-分段", 0, 0},
		{"方案7-RFC简化", 0, 0},
		{"方案8-字符串", 0, 0},
		{"方案9-组合优化", 0, 0},
		{"方案10-ASCII", 0, 0},
		{"方案12-极简", 0, 0},
	}

	validators := map[string]func(string) bool{
		"方案1-正则":        validateEmailRegex,
		"方案2-分割":        validateEmailSplit,
		"方案3-状态机":       validateEmailStateMachine,
		"方案4-LastIndex":  validateEmailLastIndex,
		"方案5-快速路径":      validateEmailFastPath,
		"方案6-分段":        validateEmailSegmented,
		"方案7-RFC简化":    validateEmailRFCSimplified,
		"方案8-字符串":        validateEmailStrings,
		"方案9-组合优化":      validateEmailCombined,
		"方案10-ASCII":    validateEmailASCII,
		"方案12-极简":       validateEmailMinimal,
	}

	allEmails := append(validEmails, invalidEmails...)

	for i, result := range results {
		validCount := 0
		for _, email := range allEmails {
			if validators[result.name](email) {
				validCount++
			}
		}
		results[i].valid = validCount
		results[i].total = len(allEmails)
	}

	fmt.Println("\n========== 邮箱验证正确性测试 ==========")
	for _, r := range results {
		fmt.Printf("%-15s: %d/%d 通过\n", r.name, r.valid, r.total)
	}
}
