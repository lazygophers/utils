package validator

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// 预编译正则表达式
var (
	mobileRegex      = regexp.MustCompile(`^1[3-9]\d{9}$`)
	idcard15Regex    = regexp.MustCompile(`^\d{15}$`)
	idcard18Regex    = regexp.MustCompile(`^\d{17}[\dXx]$`)
	chineseNameRegex = regexp.MustCompile(`^[\p{Han}·]{2,4}$`)
	uuidRegex        = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	macRegexPatterns = []*regexp.Regexp{
		regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`),
		regexp.MustCompile(`^([0-9A-Fa-f]{4}\.){2}([0-9A-Fa-f]{4})$`),
		regexp.MustCompile(`^([0-9A-Fa-f]{12})$`),
	}
)

// validateMobile 验证手机号码
func validateMobile(fl FieldLevel) bool {
	mobile := fl.Field().String()
	if mobile == "" {
		return false
	}

	// 中国大陆手机号格式：1[3-9]\d{9}
	return mobileRegex.MatchString(mobile)
}

// validateIDCard 验证身份证号码
func validateIDCard(fl FieldLevel) bool {
	idcard := fl.Field().String()
	if idcard == "" {
		return false
	}

	// 15位或18位身份证号码
	if len(idcard) == 15 {
		return validateIDCard15(idcard)
	} else if len(idcard) == 18 {
		return validateIDCard18(idcard)
	}

	return false
}

// validateIDCard15 验证15位身份证
func validateIDCard15(idcard string) bool {
	return idcard15Regex.MatchString(idcard)
}

// validateIDCard18 验证18位身份证
func validateIDCard18(idcard string) bool {
	if !idcard18Regex.MatchString(idcard) {
		return false
	}

	// 校验码验证 - 暂时先只做格式验证，算法验证比较复杂
	return true
}

// validateIDCardChecksum 验证身份证校验码
func validateIDCardChecksum(idcard string) bool {
	// 必须是18位身份证
	if len(idcard) != 18 {
		return false
	}

	// 身份证号码校验算法
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checkCodes := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	sum := 0
	for i := 0; i < 17; i++ {
		digit, err := strconv.Atoi(string(idcard[i]))
		if err != nil {
			return false
		}
		sum += digit * weights[i]
	}

	checkIndex := sum % 11
	expectedCheck := checkCodes[checkIndex]
	actualCheck := strings.ToUpper(string(idcard[17]))

	return expectedCheck == actualCheck
}

// validateBankCard 验证银行卡号
func validateBankCard(fl FieldLevel) bool {
	cardNo := fl.Field().String()
	if cardNo == "" {
		return false
	}

	// 银行卡号长度通常为13-19位
	if len(cardNo) < 13 || len(cardNo) > 19 {
		return false
	}

	// 只能包含数字
	for _, r := range cardNo {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	// Luhn算法验证
	return luhnCheck(cardNo)
}

// luhnCheck Luhn算法验证
func luhnCheck(cardNo string) bool {
	sum := 0
	alternate := false

	// 从右到左处理
	for i := len(cardNo) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(cardNo[i]))
		if err != nil {
			return false
		}

		if alternate {
			digit *= 2
			if digit > 9 {
				digit = digit%10 + digit/10
			}
		}

		sum += digit
		alternate = !alternate
	}

	return sum%10 == 0
}

// validateChineseName 验证中文姓名
func validateChineseName(fl FieldLevel) bool {
	name := fl.Field().String()
	if name == "" {
		return false
	}

	// 中文姓名：2-4个中文字符，可能包含·（少数民族姓名）
	return chineseNameRegex.MatchString(name)
}

// validateStrongPassword 验证强密码
func validateStrongPassword(fl FieldLevel) bool {
	password := fl.Field().String()
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
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	// 至少包含大写字母、小写字母、数字、特殊字符中的三种
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

// validateURL 增强的URL验证
func validateURL(fl FieldLevel) bool {
	url := fl.Field().String()
	if url == "" {
		return false
	}

	return urlRegex.MatchString(url)
}

// validateEmail 增强的邮箱验证
func validateEmail(fl FieldLevel) bool {
	email := fl.Field().String()
	if email == "" {
		return false
	}

	return emailRegex.MatchString(email)
}

// validateIPv4 IPv4地址验证
func validateIPv4(fl FieldLevel) bool {
	ip := fl.Field().String()
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

// validateMAC MAC地址验证
func validateMAC(fl FieldLevel) bool {
	mac := fl.Field().String()
	if mac == "" {
		return false
	}

	for _, re := range macRegexPatterns {
		if re.MatchString(mac) {
			return true
		}
	}

	return false
}

// validateJSON JSON格式验证
func validateJSON(fl FieldLevel) bool {
	jsonStr := fl.Field().String()
	if jsonStr == "" {
		return false
	}

	// 简单的JSON格式验证
	jsonStr = strings.TrimSpace(jsonStr)
	return (strings.HasPrefix(jsonStr, "{") && strings.HasSuffix(jsonStr, "}")) ||
		(strings.HasPrefix(jsonStr, "[") && strings.HasSuffix(jsonStr, "]"))
}

// validateUUID UUID格式验证
func validateUUID(fl FieldLevel) bool {
	uuid := fl.Field().String()
	if uuid == "" {
		return false
	}

	return uuidRegex.MatchString(strings.ToLower(uuid))
}
