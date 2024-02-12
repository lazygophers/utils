package stringx

import (
	"bytes"
	"github.com/lazygophers/log"
	"strings"
)

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
