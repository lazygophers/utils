package stringx

import (
	"github.com/lazygophers/log"
	"unicode/utf16"
)

func Utf16Len[M string | []rune | []byte](str M) int {
	s := string(str)
	log.Debugf("Utf16Len: calculating UTF-16 length for string of length %d", len(s))
	if s == "" {
		log.Debug("Utf16Len: empty string provided")
		return 0
	}
	result := len(utf16.Encode([]rune(s)))
	log.Debugf("Utf16Len: UTF-16 length is %d", result)
	return result
}
