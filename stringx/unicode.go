package stringx

import "unicode"

// allMatch 检查字符串中所有字符是否都满足给定的谓词函数
func allMatch(s string, pred func(rune) bool) bool {
	for _, c := range s {
		if !pred(c) {
			return false
		}
	}
	return true
}

// hasMatch 检查字符串中是否至少有一个字符满足给定的谓词函数
func hasMatch(s string, pred func(rune) bool) bool {
	for _, c := range s {
		if pred(c) {
			return true
		}
	}
	return false
}

// isASCIIString 检查字符串是否只包含ASCII字符
func isASCIIString(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > 127 {
			return false
		}
	}
	return true
}

func AllDigit(s string) bool {
	if s == "" {
		return true
	}
	if !isASCIIString(s) {
		return allMatch(s, unicode.IsDigit)
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func HasDigit(s string) bool {
	if !isASCIIString(s) {
		return hasMatch(s, unicode.IsDigit)
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return true
		}
	}
	return false
}

func AllLetter(s string) bool {
	return allMatch(s, unicode.IsLetter)
}

func HasLetter(s string) bool {
	return hasMatch(s, unicode.IsLetter)
}

func AllSpace(s string) bool {
	if !isASCIIString(s) {
		return allMatch(s, unicode.IsSpace)
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != ' ' && c != '\t' && c != '\n' && c != '\r' {
			return false
		}
	}
	return true
}

func HasSpace(s string) bool {
	if !isASCIIString(s) {
		return hasMatch(s, unicode.IsSpace)
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ' ' || c == '\t' || c == '\n' || c == '\r' {
			return true
		}
	}
	return false
}

func AllSymbol(s string) bool {
	return allMatch(s, unicode.IsSymbol)
}

func HasSymbol(s string) bool {
	return hasMatch(s, unicode.IsSymbol)
}

func AllMark(s string) bool {
	return allMatch(s, unicode.IsMark)
}

func HasMark(s string) bool {
	return hasMatch(s, unicode.IsMark)
}

func AllPunct(s string) bool {
	return allMatch(s, unicode.IsPunct)
}

func HasPunct(s string) bool {
	return hasMatch(s, unicode.IsPunct)
}

func AllGraphic(s string) bool {
	return allMatch(s, unicode.IsGraphic)
}

func HasGraphic(s string) bool {
	return hasMatch(s, unicode.IsGraphic)
}

func AllPrint(s string) bool {
	return allMatch(s, unicode.IsPrint)
}

func HasPrint(s string) bool {
	return hasMatch(s, unicode.IsPrint)
}

func AllControl(s string) bool {
	return allMatch(s, unicode.IsControl)
}

func HasControl(s string) bool {
	return hasMatch(s, unicode.IsControl)
}

func AllUpper(s string) bool {
	if !isASCIIString(s) {
		return allMatch(s, unicode.IsUpper)
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 'A' || s[i] > 'Z' {
			return false
		}
	}
	return true
}

func HasUpper(s string) bool {
	if !isASCIIString(s) {
		return hasMatch(s, unicode.IsUpper)
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return true
		}
	}
	return false
}

func AllLower(s string) bool {
	if !isASCIIString(s) {
		return allMatch(s, unicode.IsLower)
	}
	for i := 0; i < len(s); i++ {
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}
	return true
}

func HasLower(s string) bool {
	if !isASCIIString(s) {
		return hasMatch(s, unicode.IsLower)
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			return true
		}
	}
	return false
}

func AllTitle(s string) bool {
	return allMatch(s, unicode.IsTitle)
}

func HasTitle(s string) bool {
	return hasMatch(s, unicode.IsTitle)
}

func AllLetterOrDigit(s string) bool {
	if s == "" {
		return true
	}
	if !isASCIIString(s) {
		return allMatch(s, func(c rune) bool {
			return unicode.IsLetter(c) || unicode.IsDigit(c)
		})
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		isDigit := c >= '0' && c <= '9'
		isUpper := c >= 'A' && c <= 'Z'
		isLower := c >= 'a' && c <= 'z'
		if !isDigit && !isUpper && !isLower {
			return false
		}
	}
	return true
}

func HasLetterOrDigit(s string) bool {
	if !isASCIIString(s) {
		return hasMatch(s, func(c rune) bool {
			return unicode.IsLetter(c) || unicode.IsDigit(c)
		})
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		isDigit := c >= '0' && c <= '9'
		isUpper := c >= 'A' && c <= 'Z'
		isLower := c >= 'a' && c <= 'z'
		if isDigit || isUpper || isLower {
			return true
		}
	}
	return false
}
