package stringx

import (
	"strings"
)

// ContainsAny reports whether any of the UTF-8-encoded code points in chars are within s.
func ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune(s string, r rune) bool {
	return strings.ContainsRune(s, r)
}

// Count counts the number of non-overlapping instances of substr in s.
func Count(s, substr string) int {
	if substr == "" {
		return len([]rune(s)) + 1
	}
	return strings.Count(s, substr)
}

// EqualFold reports whether s and t are equal under Unicode case-folding.
func EqualFold(s, t string) bool {
	return strings.EqualFold(s, t)
}

// Fields splits the string s around each instance of one or more consecutive white space characters.
func Fields(s string) []string {
	return strings.Fields(s)
}

// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c).
func FieldsFunc(s string, f func(rune) bool) []string {
	if f == nil {
		return nil
	}
	return strings.FieldsFunc(s, f)
}

// HasPrefix tests whether the string s begins with prefix.
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s, substr string) int {
	return strings.Index(s, substr)
}

// IndexAny returns the index of the first instance of any Unicode code point from chars in s,
// or -1 if no Unicode code point from chars is present in s.
func IndexAny(s, chars string) int {
	return strings.IndexAny(s, chars)
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex(s, substr string) int {
	return strings.LastIndex(s, substr)
}

// LastIndexAny returns the index of the last instance of any Unicode code point from chars in s,
// or -1 if no Unicode code point from chars is present in s.
func LastIndexAny(s, chars string) int {
	return strings.LastIndexAny(s, chars)
}

// Repeat returns a new string consisting of count copies of the string s.
func Repeat(s string, count int) string {
	if count < 0 {
		return ""
	}
	return strings.Repeat(s, count)
}

// Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
func Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// ReplaceAll returns a copy of the string s with all non-overlapping instances of old replaced by new.
func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// SplitAfter slices s into all substrings after each instance of sep and returns a slice of those substrings.
func SplitAfter(s, sep string) []string {
	return strings.SplitAfter(s, sep)
}

// SplitN slices s into substrings separated by sep and returns a slice of the substrings between those separators.
func SplitN(s, sep string, n int) []string {
	return strings.SplitN(s, sep, n)
}

// SplitAfterN slices s into substrings after each instance of sep and returns a slice of those substrings.
func SplitAfterN(s, sep string, n int) []string {
	return strings.SplitAfterN(s, sep, n)
}

// Title returns a copy of the string s with all Unicode letters that begin words mapped to their Unicode title case.
func Title(s string) string {
	return strings.Title(s)
}

// ToLower returns s with all Unicode letters mapped to their lower case.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToTitle returns a copy of the string s with all Unicode letters mapped to their Unicode title case.
func ToTitle(s string) string {
	return strings.ToTitle(s)
}

// ToUpper returns s with all Unicode letters mapped to their upper case.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.
func Trim(s, cutset string) string {
	return strings.Trim(s, cutset)
}

// TrimLeft returns a slice of the string s with all leading Unicode code points contained in cutset removed.
func TrimLeft(s, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

// TrimRight returns a slice of the string s with all trailing Unicode code points contained in cutset removed.
func TrimRight(s, cutset string) string {
	return strings.TrimRight(s, cutset)
}

// TrimSpace returns a slice of the string s, with all leading and trailing white space removed.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// TrimPrefix returns s without the provided leading prefix string.
func TrimPrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

// TrimSuffix returns s without the provided trailing suffix string.
func TrimSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}
