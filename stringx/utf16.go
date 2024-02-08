package stringx

import "unicode/utf16"

func Utf16Len[M string | []rune | []byte](str M) int {
	return len(utf16.Encode([]rune(string(str))))
}
