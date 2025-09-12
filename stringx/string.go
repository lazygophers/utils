package stringx

import (
	"bytes"
	"github.com/lazygophers/log"
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

// Camel2Snake 驼峰转蛇形
func Camel2Snake(s string) string {
	var b bytes.Buffer
	for i, v := range s {
		if v >= 'A' && v <= 'Z' {
			if i > 0 {
				b.WriteString("_")
			}
			b.WriteString(string(v + 32))
		} else {
			b.WriteString(string(v))
		}
	}
	return b.String()
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

// ToSnake 蛇形
func ToSnake(s string) string {
	var b bytes.Buffer
	runes := []rune(s)
	for i, v := range runes {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			needsUnderscore := false
			
			// Check if we need an underscore before this character
			if i > 0 {
				prev := runes[i-1]
				// Add underscore for uppercase letters (camelCase -> camel_case)
				if unicode.IsUpper(v) && unicode.IsLetter(prev) {
					needsUnderscore = true
				}
				// Add underscore when transitioning from letter to number
				if unicode.IsNumber(v) && unicode.IsLetter(prev) {
					needsUnderscore = true
				}
				// Add underscore when transitioning from number to letter  
				if unicode.IsLetter(v) && unicode.IsNumber(prev) {
					needsUnderscore = true
				}
			}
			
			if needsUnderscore {
				b.WriteRune('_')
			}
			
			if unicode.IsUpper(v) {
				b.WriteRune(unicode.ToLower(v))
			} else {
				b.WriteRune(v)
			}
		} else {
			// Only add underscore if the last character wasn't an underscore
			if b.Len() > 0 {
				lastRune := []rune(b.String())
				if len(lastRune) == 0 || lastRune[len(lastRune)-1] != '_' {
					b.WriteRune('_')
				}
			}
		}
	}

	return b.String()
}

// ToKebab
func ToKebab(s string) string {
	var b bytes.Buffer
	runes := []rune(s)
	for i, v := range runes {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			needsHyphen := false
			
			// Check if we need a hyphen before this character
			if i > 0 {
				prev := runes[i-1]
				// Add hyphen for uppercase letters (camelCase -> camel-case)
				if unicode.IsUpper(v) && unicode.IsLetter(prev) {
					needsHyphen = true
				}
				// Add hyphen when transitioning from letter to number
				if unicode.IsNumber(v) && unicode.IsLetter(prev) {
					needsHyphen = true
				}
				// Add hyphen when transitioning from number to letter  
				if unicode.IsLetter(v) && unicode.IsNumber(prev) {
					needsHyphen = true
				}
			}
			
			if needsHyphen {
				b.WriteRune('-')
			}
			
			if unicode.IsUpper(v) {
				b.WriteRune(unicode.ToLower(v))
			} else {
				b.WriteRune(v)
			}
		} else {
			// Only add hyphen if the last character wasn't a hyphen
			if b.Len() > 0 {
				lastRune := []rune(b.String())
				if len(lastRune) == 0 || lastRune[len(lastRune)-1] != '-' {
					b.WriteRune('-')
				}
			}
		}
	}

	return b.String()
}

// ToCamel 转驼峰
func ToCamel(s string) string {
	var b bytes.Buffer
	upper := true
	prevWasNumber := false
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			if upper || (prevWasNumber && unicode.IsLetter(v)) {
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

// SplitLen 按长度分割字符串
func SplitLen(s string, max int) []string {
	if max <= 0 {
		return []string{s}
	}
	if s == "" {
		return []string{}
	}
	var lines []string
	b := log.GetBuffer()
	defer log.PutBuffer(b)
	
	runeCount := 0
	for _, r := range []rune(s) {
		b.WriteRune(r)
		runeCount++
		if runeCount >= max {
			lines = append(lines, b.String())
			b.Reset()
			runeCount = 0
		}
	}

	if b.Len() > 0 {
		lines = append(lines, b.String())
	}

	return lines
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
