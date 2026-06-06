package validator

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// validateAlphaSpace ASCII 字母 + 空格验证
// 零分配：字节级检查，仅允许 A-Z, a-z, 空格
func validateAlphaSpace(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || c == ' ') {
			return false
		}
	}
	return true
}

// validateAlphanumSpace ASCII 字母 + 数字 + 空格验证
// 零分配：字节级检查，仅允许 A-Z, a-z, 0-9, 空格
func validateAlphanumSpace(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == ' ') {
			return false
		}
	}
	return true
}

// validateAlphaUnicode Unicode 字母验证（含中文等多语言字母）
func validateAlphaUnicode(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// validateAlphanumUnicode Unicode 字母 + 数字验证（含中文等多语言字母）
func validateAlphanumUnicode(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// validateASCII ASCII 字符验证（所有字节 < 128）
func validateASCII(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] > 127 {
			return false
		}
	}
	return true
}

// validatePrintASCII 可打印 ASCII 字符验证（字节 32-126）
func validatePrintASCII(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] > 126 {
			return false
		}
	}
	return true
}

// validateBoolean 布尔值字符串验证（"true"/"false"/"1"/"0"）
func validateBoolean(fl FieldLevel) bool {
	s := fl.Field().String()
	return s == "true" || s == "false" || s == "1" || s == "0"
}

// validateNumber 整数验证（纯数字字符串，仅 0-9）
// 零分配：字节级检查
func validateNumber(fl FieldLevel) bool {
	s := fl.Field().String()
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// validateMultibyte 多字节字符验证（含至少一个非 ASCII 字符）
func validateMultibyte(fl FieldLevel) bool {
	s := fl.Field().String()
	for i := 0; i < len(s); i++ {
		if s[i] >= 0x80 {
			return true
		}
	}
	return false
}

// validateContains 包含子串验证
func validateContains(fl FieldLevel) bool {
	return strings.Contains(fl.Field().String(), fl.Param())
}

// validateContainsAny 包含任一字符验证
func validateContainsAny(fl FieldLevel) bool {
	return strings.ContainsAny(fl.Field().String(), fl.Param())
}

// validateContainsRune 包含指定 rune 验证
func validateContainsRune(fl FieldLevel) bool {
	param := fl.Param()
	if len(param) == 0 {
		return false
	}
	r, _ := utf8.DecodeRuneInString(param)
	return strings.ContainsRune(fl.Field().String(), r)
}

// validateStartsWith 以指定字符串开头验证
func validateStartsWith(fl FieldLevel) bool {
	return strings.HasPrefix(fl.Field().String(), fl.Param())
}

// validateStartsNotWith 不以指定字符串开头验证
func validateStartsNotWith(fl FieldLevel) bool {
	return !strings.HasPrefix(fl.Field().String(), fl.Param())
}

// validateEndsWith 以指定字符串结尾验证
func validateEndsWith(fl FieldLevel) bool {
	return strings.HasSuffix(fl.Field().String(), fl.Param())
}

// validateEndsNotWith 不以指定字符串结尾验证
func validateEndsNotWith(fl FieldLevel) bool {
	return !strings.HasSuffix(fl.Field().String(), fl.Param())
}

// validateExcludes 不包含子串验证
func validateExcludes(fl FieldLevel) bool {
	return !strings.Contains(fl.Field().String(), fl.Param())
}

// validateExcludesAll 不包含任一字符验证
func validateExcludesAll(fl FieldLevel) bool {
	return !strings.ContainsAny(fl.Field().String(), fl.Param())
}

// validateExcludesRune 不包含指定 rune 验证
func validateExcludesRune(fl FieldLevel) bool {
	param := fl.Param()
	if len(param) == 0 {
		return true
	}
	r, _ := utf8.DecodeRuneInString(param)
	return !strings.ContainsRune(fl.Field().String(), r)
}
