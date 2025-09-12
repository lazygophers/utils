package stringx

import (
	"reflect"
	"testing"
	"unicode"
)

func TestContainsAny(t *testing.T) {
	testCases := []struct {
		s        string
		chars    string
		expected bool
	}{
		{"hello", "xyz", false},
		{"hello", "el", true},
		{"hello", "", false},
		{"", "xyz", false},
		{"", "", false},
		{"测试", "试", true},
		{"测试", "abc", false},
		{"hello world", " ", true},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := ContainsAny(tc.s, tc.chars)
			if result != tc.expected {
				t.Errorf("ContainsAny(%q, %q) = %v, expected %v", tc.s, tc.chars, result, tc.expected)
			}
		})
	}
}

func TestContainsRune(t *testing.T) {
	testCases := []struct {
		s        string
		r        rune
		expected bool
	}{
		{"hello", 'h', true},
		{"hello", 'x', false},
		{"", 'a', false},
		{"测试", '测', true},
		{"测试", 'a', false},
		{"hello world", ' ', true},
		{"😀😁😂", '😀', true},
		{"😀😁😂", '😎', false},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := ContainsRune(tc.s, tc.r)
			if result != tc.expected {
				t.Errorf("ContainsRune(%q, %q) = %v, expected %v", tc.s, tc.r, result, tc.expected)
			}
		})
	}
}

func TestCount(t *testing.T) {
	testCases := []struct {
		s        string
		substr   string
		expected int
	}{
		{"hello", "l", 2},
		{"hello", "ll", 1},
		{"hello", "x", 0},
		{"", "a", 0},
		{"hello", "", 6}, // Special case: empty substring
		{"", "", 1},      // Special case: both empty
		{"测试测试", "测", 2},
		{"abcabcabc", "abc", 3},
		{"aaa", "aa", 1}, // Non-overlapping
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := Count(tc.s, tc.substr)
			if result != tc.expected {
				t.Errorf("Count(%q, %q) = %d, expected %d", tc.s, tc.substr, result, tc.expected)
			}
		})
	}
}

func TestEqualFold(t *testing.T) {
	testCases := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"hello", "HELLO", true},
		{"hello", "Hello", true},
		{"hello", "world", false},
		{"", "", true},
		{"", "a", false},
		{"测试", "测试", true},
		{"Ñoël", "ñoël", true},
		{"straße", "STRASSE", false}, // German ß doesn't fold to SS
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := EqualFold(tc.s1, tc.s2)
			if result != tc.expected {
				t.Errorf("EqualFold(%q, %q) = %v, expected %v", tc.s1, tc.s2, result, tc.expected)
			}
		})
	}
}

func TestFields(t *testing.T) {
	testCases := []struct {
		s        string
		expected []string
	}{
		{"hello world", []string{"hello", "world"}},
		{"  hello   world  ", []string{"hello", "world"}},
		{"", []string{}},
		{"   ", []string{}},
		{"single", []string{"single"}},
		{"a\tb\nc\rd", []string{"a", "b", "c", "d"}},
		{"测试 字符串", []string{"测试", "字符串"}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := Fields(tc.s)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Fields(%q) = %v, expected %v", tc.s, result, tc.expected)
			}
		})
	}
}

func TestFieldsFunc(t *testing.T) {
	t.Run("nil_function", func(t *testing.T) {
		result := FieldsFunc("hello", nil)
		if result != nil {
			t.Errorf("FieldsFunc with nil function should return nil, got %v", result)
		}
	})

	t.Run("comma_separator", func(t *testing.T) {
		f := func(r rune) bool { return r == ',' }
		result := FieldsFunc("a,b,c", f)
		expected := []string{"a", "b", "c"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("FieldsFunc comma separator = %v, expected %v", result, expected)
		}
	})

	t.Run("digit_separator", func(t *testing.T) {
		f := func(r rune) bool { return unicode.IsDigit(r) }
		result := FieldsFunc("a1b2c3d", f)
		expected := []string{"a", "b", "c", "d"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("FieldsFunc digit separator = %v, expected %v", result, expected)
		}
	})

	t.Run("empty_string", func(t *testing.T) {
		f := func(r rune) bool { return r == ',' }
		result := FieldsFunc("", f)
		expected := []string{}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("FieldsFunc empty string = %v, expected %v", result, expected)
		}
	})
}

func TestHasPrefix(t *testing.T) {
	testCases := []struct {
		s        string
		prefix   string
		expected bool
	}{
		{"hello world", "hello", true},
		{"hello world", "world", false},
		{"hello world", "", true},
		{"", "hello", false},
		{"", "", true},
		{"测试字符串", "测试", true},
		{"测试字符串", "字符", false},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := HasPrefix(tc.s, tc.prefix)
			if result != tc.expected {
				t.Errorf("HasPrefix(%q, %q) = %v, expected %v", tc.s, tc.prefix, result, tc.expected)
			}
		})
	}
}

func TestHasSuffix(t *testing.T) {
	testCases := []struct {
		s        string
		suffix   string
		expected bool
	}{
		{"hello world", "world", true},
		{"hello world", "hello", false},
		{"hello world", "", true},
		{"", "world", false},
		{"", "", true},
		{"测试字符串", "字符串", true},
		{"测试字符串", "测试", false},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := HasSuffix(tc.s, tc.suffix)
			if result != tc.expected {
				t.Errorf("HasSuffix(%q, %q) = %v, expected %v", tc.s, tc.suffix, result, tc.expected)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	testCases := []struct {
		s        string
		substr   string
		expected int
	}{
		{"hello world", "world", 6},
		{"hello world", "hello", 0},
		{"hello world", "xyz", -1},
		{"hello world", "", 0},
		{"", "hello", -1},
		{"", "", 0},
		{"测试字符串", "字符", 6}, // Note: byte index, not rune index
		{"测试字符串", "abc", -1},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := Index(tc.s, tc.substr)
			if result != tc.expected {
				t.Errorf("Index(%q, %q) = %d, expected %d", tc.s, tc.substr, result, tc.expected)
			}
		})
	}
}

func TestIndexAny(t *testing.T) {
	testCases := []struct {
		s        string
		chars    string
		expected int
	}{
		{"hello", "aeiou", 1}, // 'e' at index 1
		{"hello", "xyz", -1},
		{"hello", "", -1},
		{"", "aeiou", -1},
		{"", "", -1},
		{"测试", "试", 3}, // '试' starts at byte index 3
		{"hello world", " ", 5},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := IndexAny(tc.s, tc.chars)
			if result != tc.expected {
				t.Errorf("IndexAny(%q, %q) = %d, expected %d", tc.s, tc.chars, result, tc.expected)
			}
		})
	}
}

func TestLastIndex(t *testing.T) {
	testCases := []struct {
		s        string
		substr   string
		expected int
	}{
		{"hello world hello", "hello", 12},
		{"hello world hello", "world", 6},
		{"hello world hello", "xyz", -1},
		{"hello world hello", "", 17},
		{"", "hello", -1},
		{"", "", 0},
		{"测试测试", "测试", 6}, // Last occurrence
		{"abcabc", "abc", 3},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := LastIndex(tc.s, tc.substr)
			if result != tc.expected {
				t.Errorf("LastIndex(%q, %q) = %d, expected %d", tc.s, tc.substr, result, tc.expected)
			}
		})
	}
}

func TestLastIndexAny(t *testing.T) {
	testCases := []struct {
		s        string
		chars    string
		expected int
	}{
		{"hello", "aeiou", 4}, // Last vowel 'o' at index 4
		{"hello world", "aeiou", 7}, // 'o' at index 7
		{"hello", "xyz", -1},
		{"hello", "", -1},
		{"", "aeiou", -1},
		{"测试字符串", "试串", 12}, // '串' is the last match
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := LastIndexAny(tc.s, tc.chars)
			if result != tc.expected {
				t.Errorf("LastIndexAny(%q, %q) = %d, expected %d", tc.s, tc.chars, result, tc.expected)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	testCases := []struct {
		s        string
		count    int
		expected string
	}{
		{"a", 3, "aaa"},
		{"hello", 2, "hellohello"},
		{"", 5, ""},
		{"test", 0, ""},
		{"test", -1, ""}, // Negative count
		{"测试", 2, "测试测试"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := Repeat(tc.s, tc.count)
			if result != tc.expected {
				t.Errorf("Repeat(%q, %d) = %q, expected %q", tc.s, tc.count, result, tc.expected)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	testCases := []struct {
		s        string
		old      string
		new      string
		n        int
		expected string
	}{
		{"hello world hello", "hello", "hi", 1, "hi world hello"},
		{"hello world hello", "hello", "hi", 2, "hi world hi"},
		{"hello world hello", "hello", "hi", -1, "hi world hi"},
		{"hello world hello", "xyz", "abc", 1, "hello world hello"},
		{"", "old", "new", 1, ""},
		{"test", "", "x", 3, "xtxexst"}, // Replace empty string (limited by n)
		{"测试测试", "测试", "检查", 1, "检查测试"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := Replace(tc.s, tc.old, tc.new, tc.n)
			if result != tc.expected {
				t.Errorf("Replace(%q, %q, %q, %d) = %q, expected %q", tc.s, tc.old, tc.new, tc.n, result, tc.expected)
			}
		})
	}
}

func TestReplaceAll(t *testing.T) {
	testCases := []struct {
		s        string
		old      string
		new      string
		expected string
	}{
		{"hello world hello", "hello", "hi", "hi world hi"},
		{"hello world hello", "xyz", "abc", "hello world hello"},
		{"", "old", "new", ""},
		{"aaa", "aa", "b", "ba"}, // Non-overlapping replacement
		{"测试测试测试", "测试", "检查", "检查检查检查"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := ReplaceAll(tc.s, tc.old, tc.new)
			if result != tc.expected {
				t.Errorf("ReplaceAll(%q, %q, %q) = %q, expected %q", tc.s, tc.old, tc.new, result, tc.expected)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	testCases := []struct {
		s        string
		sep      string
		expected []string
	}{
		{"a,b,c", ",", []string{"a", "b", "c"}},
		{"a,,b", ",", []string{"a", "", "b"}},
		{"", ",", []string{""}},
		{"abc", ",", []string{"abc"}},
		{"a,b,c", "", []string{"a", ",", "b", ",", "c"}}, // Empty separator
		{"测试,字符串", ",", []string{"测试", "字符串"}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := Split(tc.s, tc.sep)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Split(%q, %q) = %v, expected %v", tc.s, tc.sep, result, tc.expected)
			}
		})
	}
}

func TestSplitAfter(t *testing.T) {
	testCases := []struct {
		s        string
		sep      string
		expected []string
	}{
		{"a,b,c", ",", []string{"a,", "b,", "c"}},
		{"a,,b", ",", []string{"a,", ",", "b"}},
		{"", ",", []string{""}},
		{"abc", ",", []string{"abc"}},
		{"测试,字符串", ",", []string{"测试,", "字符串"}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := SplitAfter(tc.s, tc.sep)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("SplitAfter(%q, %q) = %v, expected %v", tc.s, tc.sep, result, tc.expected)
			}
		})
	}
}

func TestSplitN(t *testing.T) {
	testCases := []struct {
		s        string
		sep      string
		n        int
		expected []string
	}{
		{"a,b,c,d", ",", 2, []string{"a", "b,c,d"}},
		{"a,b,c,d", ",", 0, nil},
		{"a,b,c,d", ",", -1, []string{"a", "b", "c", "d"}},
		{"", ",", 2, []string{""}},
		{"abc", ",", 2, []string{"abc"}},
		{"测试,字符串,测试", ",", 2, []string{"测试", "字符串,测试"}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := SplitN(tc.s, tc.sep, tc.n)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("SplitN(%q, %q, %d) = %v, expected %v", tc.s, tc.sep, tc.n, result, tc.expected)
			}
		})
	}
}

func TestSplitAfterN(t *testing.T) {
	testCases := []struct {
		s        string
		sep      string
		n        int
		expected []string
	}{
		{"a,b,c,d", ",", 2, []string{"a,", "b,c,d"}},
		{"a,b,c,d", ",", 0, nil},
		{"a,b,c,d", ",", -1, []string{"a,", "b,", "c,", "d"}},
		{"", ",", 2, []string{""}},
		{"abc", ",", 2, []string{"abc"}},
		{"测试,字符串,测试", ",", 2, []string{"测试,", "字符串,测试"}},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := SplitAfterN(tc.s, tc.sep, tc.n)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("SplitAfterN(%q, %q, %d) = %v, expected %v", tc.s, tc.sep, tc.n, result, tc.expected)
			}
		})
	}
}

func TestTitle(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello world", "Hello World"},
		{"", ""},
		{"a", "A"},
		{"123abc", "123abc"}, // Numbers don't create word boundaries
		{"测试 字符串", "测试 字符串"}, // Chinese characters don't change
		{"hello-world", "Hello-World"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := Title(tc.input)
			if result != tc.expected {
				t.Errorf("Title(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"HELLO", "hello"},
		{"Hello World", "hello world"},
		{"", ""},
		{"123ABC", "123abc"},
		{"测试", "测试"}, // Chinese characters don't change
		{"Ñoël", "ñoël"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := ToLower(tc.input)
			if result != tc.expected {
				t.Errorf("ToLower(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestToTitle(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello world", "HELLO WORLD"},
		{"", ""},
		{"Hello World", "HELLO WORLD"},
		{"123abc", "123ABC"},
		{"测试", "测试"}, // Chinese characters don't change
		{"ñoël", "ÑOËL"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := ToTitle(tc.input)
			if result != tc.expected {
				t.Errorf("ToTitle(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "HELLO"},
		{"Hello World", "HELLO WORLD"},
		{"", ""},
		{"123abc", "123ABC"},
		{"测试", "测试"}, // Chinese characters don't change
		{"ñoël", "ÑOËL"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := ToUpper(tc.input)
			if result != tc.expected {
				t.Errorf("ToUpper(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	testCases := []struct {
		s        string
		cutset   string
		expected string
	}{
		{"  hello  ", " ", "hello"},
		{"!!!hello!!!", "!", "hello"},
		{"", " ", ""},
		{"hello", "", "hello"},
		{"abcdefabc", "abc", "def"},
		{"测试字符串测试", "测试", "字符串"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := Trim(tc.s, tc.cutset)
			if result != tc.expected {
				t.Errorf("Trim(%q, %q) = %q, expected %q", tc.s, tc.cutset, result, tc.expected)
			}
		})
	}
}

func TestTrimLeft(t *testing.T) {
	testCases := []struct {
		s        string
		cutset   string
		expected string
	}{
		{"  hello  ", " ", "hello  "},
		{"!!!hello!!!", "!", "hello!!!"},
		{"", " ", ""},
		{"hello", "", "hello"},
		{"abcdefabc", "abc", "defabc"},
		{"测试字符串测试", "测试", "字符串测试"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := TrimLeft(tc.s, tc.cutset)
			if result != tc.expected {
				t.Errorf("TrimLeft(%q, %q) = %q, expected %q", tc.s, tc.cutset, result, tc.expected)
			}
		})
	}
}

func TestTrimRight(t *testing.T) {
	testCases := []struct {
		s        string
		cutset   string
		expected string
	}{
		{"  hello  ", " ", "  hello"},
		{"!!!hello!!!", "!", "!!!hello"},
		{"", " ", ""},
		{"hello", "", "hello"},
		{"abcdefabc", "abc", "abcdef"},
		{"测试字符串测试", "测试", "测试字符串"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := TrimRight(tc.s, tc.cutset)
			if result != tc.expected {
				t.Errorf("TrimRight(%q, %q) = %q, expected %q", tc.s, tc.cutset, result, tc.expected)
			}
		})
	}
}

func TestTrimSpace(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"  hello  ", "hello"},
		{"\t\nhello\r\n", "hello"},
		{"", ""},
		{"   ", ""},
		{"hello", "hello"},
		{" 测试 ", "测试"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := TrimSpace(tc.input)
			if result != tc.expected {
				t.Errorf("TrimSpace(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestTrimPrefix(t *testing.T) {
	testCases := []struct {
		s        string
		prefix   string
		expected string
	}{
		{"hello world", "hello ", "world"},
		{"hello world", "world", "hello world"}, // Prefix not found
		{"hello world", "", "hello world"},
		{"", "prefix", ""},
		{"测试字符串", "测试", "字符串"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := TrimPrefix(tc.s, tc.prefix)
			if result != tc.expected {
				t.Errorf("TrimPrefix(%q, %q) = %q, expected %q", tc.s, tc.prefix, result, tc.expected)
			}
		})
	}
}

func TestTrimSuffix(t *testing.T) {
	testCases := []struct {
		s        string
		suffix   string
		expected string
	}{
		{"hello world", " world", "hello"},
		{"hello world", "hello", "hello world"}, // Suffix not found
		{"hello world", "", "hello world"},
		{"", "suffix", ""},
		{"测试字符串", "字符串", "测试"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := TrimSuffix(tc.s, tc.suffix)
			if result != tc.expected {
				t.Errorf("TrimSuffix(%q, %q) = %q, expected %q", tc.s, tc.suffix, result, tc.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkContainsAny(b *testing.B) {
	s := "hello world test string for benchmarking"
	chars := "xyz"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsAny(s, chars)
	}
}

func BenchmarkSplit(b *testing.B) {
	s := "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z"
	sep := ","
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Split(s, sep)
	}
}

func BenchmarkReplaceAll(b *testing.B) {
	s := "hello world hello world hello world hello world"
	old := "hello"
	new := "hi"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReplaceAll(s, old, new)
	}
}