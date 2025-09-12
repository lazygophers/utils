package stringx

import (
	"github.com/lazygophers/log"
	"strings"
)

// ContainsAny reports whether any of the UTF-8-encoded code points in chars are within s.
func ContainsAny(s, chars string) bool {
	log.Debugf("ContainsAny: checking if string %q contains any of %q", s, chars)
	result := strings.ContainsAny(s, chars)
	log.Debugf("ContainsAny result: %v", result)
	return result
}

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune(s string, r rune) bool {
	log.Debugf("ContainsRune: checking if string %q contains rune %q", s, r)
	result := strings.ContainsRune(s, r)
	log.Debugf("ContainsRune result: %v", result)
	return result
}

// Count counts the number of non-overlapping instances of substr in s.
func Count(s, substr string) int {
	log.Debugf("Count: counting occurrences of %q in %q", substr, s)
	if substr == "" {
		log.Warn("Count: empty substring provided")
		return len([]rune(s)) + 1
	}
	result := strings.Count(s, substr)
	log.Debugf("Count result: %d", result)
	return result
}

// EqualFold reports whether s and t are equal under Unicode case-folding.
func EqualFold(s, t string) bool {
	log.Debugf("EqualFold: comparing %q and %q", s, t)
	result := strings.EqualFold(s, t)
	log.Debugf("EqualFold result: %v", result)
	return result
}

// Fields splits the string s around each instance of one or more consecutive white space characters.
func Fields(s string) []string {
	log.Debugf("Fields: splitting string %q by whitespace", s)
	result := strings.Fields(s)
	log.Debugf("Fields result: %d fields", len(result))
	return result
}

// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c).
func FieldsFunc(s string, f func(rune) bool) []string {
	log.Debugf("FieldsFunc: splitting string %q with custom function", s)
	if f == nil {
		log.Error("FieldsFunc: function parameter is nil")
		return nil
	}
	result := strings.FieldsFunc(s, f)
	log.Debugf("FieldsFunc result: %d fields", len(result))
	return result
}

// HasPrefix tests whether the string s begins with prefix.
func HasPrefix(s, prefix string) bool {
	log.Debugf("HasPrefix: checking if %q has prefix %q", s, prefix)
	result := strings.HasPrefix(s, prefix)
	log.Debugf("HasPrefix result: %v", result)
	return result
}

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool {
	log.Debugf("HasSuffix: checking if %q has suffix %q", s, suffix)
	result := strings.HasSuffix(s, suffix)
	log.Debugf("HasSuffix result: %v", result)
	return result
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s, substr string) int {
	log.Debugf("Index: finding first occurrence of %q in %q", substr, s)
	result := strings.Index(s, substr)
	log.Debugf("Index result: %d", result)
	return result
}

// IndexAny returns the index of the first instance of any Unicode code point from chars in s,
// or -1 if no Unicode code point from chars is present in s.
func IndexAny(s, chars string) int {
	log.Debugf("IndexAny: finding first occurrence of any char from %q in %q", chars, s)
	result := strings.IndexAny(s, chars)
	log.Debugf("IndexAny result: %d", result)
	return result
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex(s, substr string) int {
	log.Debugf("LastIndex: finding last occurrence of %q in %q", substr, s)
	result := strings.LastIndex(s, substr)
	log.Debugf("LastIndex result: %d", result)
	return result
}

// LastIndexAny returns the index of the last instance of any Unicode code point from chars in s,
// or -1 if no Unicode code point from chars is present in s.
func LastIndexAny(s, chars string) int {
	log.Debugf("LastIndexAny: finding last occurrence of any char from %q in %q", chars, s)
	result := strings.LastIndexAny(s, chars)
	log.Debugf("LastIndexAny result: %d", result)
	return result
}

// Repeat returns a new string consisting of count copies of the string s.
func Repeat(s string, count int) string {
	log.Debugf("Repeat: repeating %q %d times", s, count)
	if count < 0 {
		log.Error("Repeat: negative count provided")
		return ""
	}
	result := strings.Repeat(s, count)
	log.Debugf("Repeat result length: %d", len(result))
	return result
}

// Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
func Replace(s, old, new string, n int) string {
	log.Debugf("Replace: replacing %q with %q in %q (max %d times)", old, new, s, n)
	result := strings.Replace(s, old, new, n)
	log.Debugf("Replace completed")
	return result
}

// ReplaceAll returns a copy of the string s with all non-overlapping instances of old replaced by new.
func ReplaceAll(s, old, new string) string {
	log.Debugf("ReplaceAll: replacing all %q with %q in %q", old, new, s)
	result := strings.ReplaceAll(s, old, new)
	log.Debugf("ReplaceAll completed")
	return result
}

// Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
func Split(s, sep string) []string {
	log.Debugf("Split: splitting %q by separator %q", s, sep)
	result := strings.Split(s, sep)
	log.Debugf("Split result: %d parts", len(result))
	return result
}

// SplitAfter slices s into all substrings after each instance of sep and returns a slice of those substrings.
func SplitAfter(s, sep string) []string {
	log.Debugf("SplitAfter: splitting %q after separator %q", s, sep)
	result := strings.SplitAfter(s, sep)
	log.Debugf("SplitAfter result: %d parts", len(result))
	return result
}

// SplitN slices s into substrings separated by sep and returns a slice of the substrings between those separators.
func SplitN(s, sep string, n int) []string {
	log.Debugf("SplitN: splitting %q by separator %q (max %d parts)", s, sep, n)
	result := strings.SplitN(s, sep, n)
	log.Debugf("SplitN result: %d parts", len(result))
	return result
}

// SplitAfterN slices s into substrings after each instance of sep and returns a slice of those substrings.
func SplitAfterN(s, sep string, n int) []string {
	log.Debugf("SplitAfterN: splitting %q after separator %q (max %d parts)", s, sep, n)
	result := strings.SplitAfterN(s, sep, n)
	log.Debugf("SplitAfterN result: %d parts", len(result))
	return result
}

// Title returns a copy of the string s with all Unicode letters that begin words mapped to their Unicode title case.
func Title(s string) string {
	log.Debugf("Title: converting %q to title case", s)
	result := strings.Title(s)
	log.Debugf("Title completed")
	return result
}

// ToLower returns s with all Unicode letters mapped to their lower case.
func ToLower(s string) string {
	log.Debugf("ToLower: converting %q to lowercase", s)
	result := strings.ToLower(s)
	log.Debugf("ToLower completed")
	return result
}

// ToTitle returns a copy of the string s with all Unicode letters mapped to their Unicode title case.
func ToTitle(s string) string {
	log.Debugf("ToTitle: converting %q to title case", s)
	result := strings.ToTitle(s)
	log.Debugf("ToTitle completed")
	return result
}

// ToUpper returns s with all Unicode letters mapped to their upper case.
func ToUpper(s string) string {
	log.Debugf("ToUpper: converting %q to uppercase", s)
	result := strings.ToUpper(s)
	log.Debugf("ToUpper completed")
	return result
}

// Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.
func Trim(s, cutset string) string {
	log.Debugf("Trim: trimming %q with cutset %q", s, cutset)
	result := strings.Trim(s, cutset)
	log.Debugf("Trim completed")
	return result
}

// TrimLeft returns a slice of the string s with all leading Unicode code points contained in cutset removed.
func TrimLeft(s, cutset string) string {
	log.Debugf("TrimLeft: trimming left of %q with cutset %q", s, cutset)
	result := strings.TrimLeft(s, cutset)
	log.Debugf("TrimLeft completed")
	return result
}

// TrimRight returns a slice of the string s with all trailing Unicode code points contained in cutset removed.
func TrimRight(s, cutset string) string {
	log.Debugf("TrimRight: trimming right of %q with cutset %q", s, cutset)
	result := strings.TrimRight(s, cutset)
	log.Debugf("TrimRight completed")
	return result
}

// TrimSpace returns a slice of the string s, with all leading and trailing white space removed.
func TrimSpace(s string) string {
	log.Debugf("TrimSpace: trimming whitespace from %q", s)
	result := strings.TrimSpace(s)
	log.Debugf("TrimSpace completed")
	return result
}

// TrimPrefix returns s without the provided leading prefix string.
func TrimPrefix(s, prefix string) string {
	log.Debugf("TrimPrefix: removing prefix %q from %q", prefix, s)
	result := strings.TrimPrefix(s, prefix)
	log.Debugf("TrimPrefix completed")
	return result
}

// TrimSuffix returns s without the provided trailing suffix string.
func TrimSuffix(s, suffix string) string {
	log.Debugf("TrimSuffix: removing suffix %q from %q", suffix, s)
	result := strings.TrimSuffix(s, suffix)
	log.Debugf("TrimSuffix completed")
	return result
}
