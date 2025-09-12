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
	log.Debugf("ToString: converting byte slice of length %d", len(b))
	if b == nil {
		log.Debug("ToString: nil byte slice provided")
		return ""
	}
	if len(b) == 0 {
		log.Debug("ToString: empty byte slice provided")
		return ""
	}
	result := *(*string)(unsafe.Pointer(&b))
	log.Debugf("ToString: converted to string of length %d", len(result))
	return result
}

func ToBytes(s string) []byte {
	log.Debugf("ToBytes: converting string of length %d", len(s))
	if s == "" {
		log.Debug("ToBytes: empty string provided")
		return nil
	}
	result := *(*[]byte)(unsafe.Pointer(&s))
	log.Debugf("ToBytes: converted to byte slice of length %d", len(result))
	return result
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
	log.Debugf("Snake2Camel: converting %q", s)
	if s == "" {
		log.Debug("Snake2Camel: empty string provided")
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
	result := b.String()
	log.Debugf("Snake2Camel result: %q", result)
	return result
}

// Snake2SmallCamel 蛇形转小驼峰
func Snake2SmallCamel(s string) string {
	log.Debugf("Snake2SmallCamel: converting %q", s)
	if s == "" {
		log.Debug("Snake2SmallCamel: empty string provided")
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
			} else if upper {
				b.WriteRune(unicode.ToUpper(v))
				upper = false
			} else {
				b.WriteRune(v)
			}
		}
	}
	result := b.String()
	log.Debugf("Snake2SmallCamel result: %q", result)
	return result
}

// ToSnake 蛇形
func ToSnake(s string) string {
	var b bytes.Buffer
	for i, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			if unicode.IsUpper(v) {
				if i > 0 {
					b.WriteRune('_')
				}
				b.WriteRune(unicode.ToLower(v))
			} else {
				b.WriteRune(v)
			}
		} else {
			b.WriteRune('_')
		}
	}

	return b.String()
}

// ToKebab
func ToKebab(s string) string {
	var b bytes.Buffer
	for i, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			if unicode.IsUpper(v) {
				if i > 0 {
					b.WriteRune('-')
				}
				b.WriteRune(unicode.ToLower(v))
			} else {
				b.WriteRune(v)
			}
		} else {
			b.WriteRune('-')
		}
	}

	return b.String()
}

// ToCamel 转驼峰
func ToCamel(s string) string {
	var b bytes.Buffer
	upper := true
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			if upper {
				b.WriteRune(unicode.ToUpper(v))
				upper = false
			} else {
				b.WriteRune(v)
			}
		} else {
			upper = true
		}
	}
	return b.String()
}

func ToSlash(s string) string {
	var b bytes.Buffer
	for i, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			if unicode.IsUpper(v) {
				if i > 0 {
					b.WriteRune('/')
				}
				b.WriteRune(unicode.ToLower(v))
			} else {
				b.WriteRune(v)
			}
		} else {
			b.WriteRune('/')
		}
	}
	return b.String()
}

func ToDot(s string) string {
	var b bytes.Buffer
	for i, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			if unicode.IsUpper(v) {
				if i > 0 {
					b.WriteRune('.')
				}
				b.WriteRune(unicode.ToLower(v))
			} else {
				b.WriteRune(v)
			}
		} else {
			b.WriteRune('.')
		}
	}
	return b.String()
}

// ToSmallCamel 转小驼峰
func ToSmallCamel(s string) string {
	var b bytes.Buffer
	upper := false
	isFirst := true
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			if isFirst {
				isFirst = false
				b.WriteRune(unicode.ToLower(v))
			} else if upper {
				b.WriteRune(unicode.ToUpper(v))
				upper = false
			} else {
				b.WriteRune(v)
			}
		} else if !isFirst {
			upper = true
		}
	}
	return b.String()
}

// SplitLen 按长度分割字符串
func SplitLen(s string, max int) []string {
	log.Debugf("SplitLen: splitting string of length %d with max length %d", len(s), max)
	if max <= 0 {
		log.Error("SplitLen: max length must be positive")
		return []string{s}
	}
	if s == "" {
		log.Debug("SplitLen: empty string provided")
		return []string{}
	}
	var lines []string
	b := log.GetBuffer()
	defer log.PutBuffer(b)

	for _, r := range []rune(s) {
		b.WriteRune(r)
		if b.Len() >= max {
			lines = append(lines, b.String())
			b.Reset()
		}
	}

	if b.Len() > 0 {
		lines = append(lines, b.String())
	}

	log.Debugf("SplitLen result: %d lines", len(lines))
	return lines
}

// Shorten 缩短字符串
func Shorten(s string, max int) string {
	log.Debugf("Shorten: shortening string of length %d to max %d", len(s), max)
	if max < 0 {
		log.Error("Shorten: max length cannot be negative")
		return ""
	}
	if len(s) <= max {
		log.Debug("Shorten: string already within limit")
		return s
	}

	result := s[:max]
	log.Debugf("Shorten result: %q", result)
	return result
}

func ShortenShow(s string, max int) string {
	log.Debugf("ShortenShow: shortening string of length %d to max %d with ellipsis", len(s), max)
	if max < 0 {
		log.Error("ShortenShow: max length cannot be negative")
		return "..."
	}
	if len(s) <= max {
		log.Debug("ShortenShow: string already within limit")
		return s
	}

	if max < 3 {
		log.Warn("ShortenShow: max length is less than 3, returning ellipsis only")
		return "..."
	}

	result := s[:max-3] + "..."
	log.Debugf("ShortenShow result: %q", result)
	return result
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
	log.Debugf("Reverse: reversing string of length %d", len(s))
	if s == "" {
		log.Debug("Reverse: empty string provided")
		return ""
	}
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	result := string(runes)
	log.Debugf("Reverse result: %q", result)
	return result
}

func Quote(s string) string {
	return strconv.Quote(s)
}

func QuotePure(s string) string {
	return strings.TrimPrefix(strings.TrimSuffix(Quote(s), `"`), `"`)
}
