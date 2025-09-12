package stringx

import "unicode/utf16"

func Utf16Len[M string | []rune | []byte](str M) int {
	s := string(str)
	if s == "" {
		return 0
	}
	return len(utf16.Encode([]rune(s)))
}
