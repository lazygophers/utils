package stringx

import (
	"bytes"
	"github.com/lazygophers/log"
	"strings"
	"unicode"
	"unsafe"
)

func ToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ToBytes(s string) []byte {
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
	var b bytes.Buffer
	upper := true
	for _, v := range s {
		if v == '_' {
			upper = true
		} else {
			if upper {
				b.WriteString(string(v - 32))
				upper = false
			} else {
				b.WriteString(string(v))
			}
		}
	}
	return b.String()
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
	for idx, v := range s {
		if idx == 0 {
			b.WriteRune(unicode.ToLower(v))
			continue
		}

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

// SplitLen 按长度分割字符串
func SplitLen(s string, max int) []string {
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

	return lines
}

// Shorten 缩短字符串
func Shorten(s string, max int) string {
	if len(s) <= max {
		return s
	}

	return s[:max]
}

func ShortenShow(s string, max int) string {
	if len(s) <= max {
		return s
	}

	return s[:max] + "..."
}

func IsUpper[M string | rune](r M) bool {
	return string(r) == strings.ToUpper(string(r))
}

func Reverse(s string) string {
	var b bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteString(string(s[i]))
	}
	return b.String()
}
