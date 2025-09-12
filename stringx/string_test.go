package stringx

import (
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestToString(t *testing.T) {
	t.Run("nil_byte_slice", func(t *testing.T) {
		result := ToString(nil)
		if result != "" {
			t.Errorf("ToString(nil) = %q, expected empty string", result)
		}
	})

	t.Run("empty_byte_slice", func(t *testing.T) {
		result := ToString([]byte{})
		if result != "" {
			t.Errorf("ToString([]byte{}) = %q, expected empty string", result)
		}
	})

	t.Run("valid_byte_slice", func(t *testing.T) {
		input := []byte("hello world")
		result := ToString(input)
		expected := "hello world"
		if result != expected {
			t.Errorf("ToString(%v) = %q, expected %q", input, result, expected)
		}
	})

	t.Run("unicode_byte_slice", func(t *testing.T) {
		input := []byte("你好世界")
		result := ToString(input)
		expected := "你好世界"
		if result != expected {
			t.Errorf("ToString(%v) = %q, expected %q", input, result, expected)
		}
	})

	t.Run("zero_copy_verification", func(t *testing.T) {
		input := []byte("test string")
		result := ToString(input)
		
		// Verify zero-copy by comparing underlying pointers
		if len(input) > 0 && len(result) > 0 {
			inputPtr := (*reflect.StringHeader)(unsafe.Pointer(&result)).Data
			expectedPtr := (*reflect.SliceHeader)(unsafe.Pointer(&input)).Data
			if inputPtr != expectedPtr {
				t.Error("ToString should perform zero-copy conversion")
			}
		}
	})
}

func TestToBytes(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		result := ToBytes("")
		if result != nil {
			t.Errorf("ToBytes(\"\") = %v, expected nil", result)
		}
	})

	t.Run("valid_string", func(t *testing.T) {
		input := "hello world"
		result := ToBytes(input)
		expected := []byte("hello world")
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToBytes(%q) = %v, expected %v", input, result, expected)
		}
	})

	t.Run("unicode_string", func(t *testing.T) {
		input := "你好世界"
		result := ToBytes(input)
		expected := []byte("你好世界")
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ToBytes(%q) = %v, expected %v", input, result, expected)
		}
	})

	t.Run("zero_copy_verification", func(t *testing.T) {
		input := "test string"
		result := ToBytes(input)
		
		// Verify zero-copy by comparing underlying pointers
		if len(input) > 0 && len(result) > 0 {
			stringPtr := (*reflect.StringHeader)(unsafe.Pointer(&input)).Data
			slicePtr := (*reflect.SliceHeader)(unsafe.Pointer(&result)).Data
			if stringPtr != slicePtr {
				t.Error("ToBytes should perform zero-copy conversion")
			}
		}
	})
}

func TestCamel2Snake(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"a", "a"},
		{"A", "a"},
		{"CamelCase", "camel_case"},
		{"HTTPSConnection", "h_t_t_p_s_connection"},
		{"XMLParser", "x_m_l_parser"},
		{"iPhone", "i_phone"},
		{"iOS", "i_o_s"},
		{"myVariableName", "my_variable_name"},
		{"SimpleTest", "simple_test"},
		{"ABC", "a_b_c"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := Camel2Snake(tc.input)
			if result != tc.expected {
				t.Errorf("Camel2Snake(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestSnake2Camel(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"a", "A"},
		{"simple_test", "SimpleTest"},
		{"my_variable_name", "MyVariableName"},
		{"http_connection", "HttpConnection"},
		{"xml_parser", "XmlParser"},
		{"_leading_underscore", "LeadingUnderscore"},
		{"trailing_underscore_", "TrailingUnderscore"},
		{"multiple___underscores", "MultipleUnderscores"},
		{"single", "Single"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := Snake2Camel(tc.input)
			if result != tc.expected {
				t.Errorf("Snake2Camel(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestSnake2SmallCamel(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"a", "a"},
		{"simple_test", "simpleTest"},
		{"my_variable_name", "myVariableName"},
		{"http_connection", "httpConnection"},
		{"xml_parser", "xmlParser"},
		{"_leading_underscore", "leadingUnderscore"},
		{"trailing_underscore_", "trailingUnderscore"},
		{"multiple___underscores", "multipleUnderscores"},
		{"single", "single"},
		{"UPPER_CASE", "upperCase"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := Snake2SmallCamel(tc.input)
			if result != tc.expected {
				t.Errorf("Snake2SmallCamel(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestToSnake(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"simple", "simple"},
		{"SimpleTest", "simple_test"},
		{"HTTPSConnection", "h_t_t_p_s_connection"},
		{"XMLHttpParser", "x_m_l_http_parser"},
		{"iPhone15Pro", "i_phone_15_pro"},
		{"iOS16", "i_o_s_16"},
		{"myVariable123Name", "my_variable_123_name"},
		{"Test@Symbol#Here", "test_symbol_here"},
		{"Multiple   Spaces", "multiple_spaces"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ToSnake(tc.input)
			if result != tc.expected {
				t.Errorf("ToSnake(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestToKebab(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"simple", "simple"},
		{"SimpleTest", "simple-test"},
		{"HTTPSConnection", "h-t-t-p-s-connection"},
		{"XMLHttpParser", "x-m-l-http-parser"},
		{"iPhone15Pro", "i-phone-15-pro"},
		{"iOS16", "i-o-s-16"},
		{"myVariable123Name", "my-variable-123-name"},
		{"Test@Symbol#Here", "test-symbol-here"},
		{"Multiple   Spaces", "multiple-spaces"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ToKebab(tc.input)
			if result != tc.expected {
				t.Errorf("ToKebab(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestToCamel(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"simple", "Simple"},
		{"simple_test", "SimpleTest"},
		{"my-variable-name", "MyVariableName"},
		{"http.connection", "HttpConnection"},
		{"xml/parser", "XmlParser"},
		{"multiple   spaces", "MultipleSpaces"},
		{"123numbers", "123Numbers"},
		{"@symbol#test", "SymbolTest"},
		{"_leading_underscore", "LeadingUnderscore"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ToCamel(tc.input)
			if result != tc.expected {
				t.Errorf("ToCamel(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestToSlash(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"simple", "simple"},
		{"SimpleTest", "simple/test"},
		{"HTTPSConnection", "h/t/t/p/s/connection"},
		{"XMLHttpParser", "x/m/l/http/parser"},
		{"iPhone15Pro", "i/phone/15/pro"},
		{"Test@Symbol#Here", "test/symbol/here"},
		{"Multiple   Spaces", "multiple/spaces"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ToSlash(tc.input)
			if result != tc.expected {
				t.Errorf("ToSlash(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestToDot(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"simple", "simple"},
		{"SimpleTest", "simple.test"},
		{"HTTPSConnection", "h.t.t.p.s.connection"},
		{"XMLHttpParser", "x.m.l.http.parser"},
		{"iPhone15Pro", "i.phone.15.pro"},
		{"Test@Symbol#Here", "test.symbol.here"},
		{"Multiple   Spaces", "multiple.spaces"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ToDot(tc.input)
			if result != tc.expected {
				t.Errorf("ToDot(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestToSmallCamel(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"simple", "simple"},
		{"simple_test", "simpleTest"},
		{"my-variable-name", "myVariableName"},
		{"http.connection", "httpConnection"},
		{"xml/parser", "xmlParser"},
		{"multiple   spaces", "multipleSpaces"},
		{"123numbers", "123Numbers"},
		{"@symbol#test", "symbolTest"},
		{"_leading_underscore", "leadingUnderscore"},
		{"UPPER_CASE", "upperCase"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ToSmallCamel(tc.input)
			if result != tc.expected {
				t.Errorf("ToSmallCamel(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestSplitLen(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		result := SplitLen("", 5)
		expected := []string{}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SplitLen(\"\", 5) = %v, expected %v", result, expected)
		}
	})

	t.Run("negative_max", func(t *testing.T) {
		result := SplitLen("hello", -1)
		expected := []string{"hello"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SplitLen(\"hello\", -1) = %v, expected %v", result, expected)
		}
	})

	t.Run("zero_max", func(t *testing.T) {
		result := SplitLen("hello", 0)
		expected := []string{"hello"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SplitLen(\"hello\", 0) = %v, expected %v", result, expected)
		}
	})

	t.Run("normal_split", func(t *testing.T) {
		result := SplitLen("hello world test", 5)
		expected := []string{"hello", " worl", "d tes", "t"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SplitLen(\"hello world test\", 5) = %v, expected %v", result, expected)
		}
	})

	t.Run("unicode_split", func(t *testing.T) {
		result := SplitLen("你好世界测试", 2)
		expected := []string{"你好", "世界", "测试"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SplitLen(\"你好世界测试\", 2) = %v, expected %v", result, expected)
		}
	})

	t.Run("exact_division", func(t *testing.T) {
		result := SplitLen("abcdef", 3)
		expected := []string{"abc", "def"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("SplitLen(\"abcdef\", 3) = %v, expected %v", result, expected)
		}
	})
}

func TestShorten(t *testing.T) {
	t.Run("negative_max", func(t *testing.T) {
		result := Shorten("hello", -1)
		if result != "" {
			t.Errorf("Shorten(\"hello\", -1) = %q, expected empty string", result)
		}
	})

	t.Run("zero_max", func(t *testing.T) {
		result := Shorten("hello", 0)
		if result != "" {
			t.Errorf("Shorten(\"hello\", 0) = %q, expected empty string", result)
		}
	})

	t.Run("string_shorter_than_max", func(t *testing.T) {
		input := "hello"
		result := Shorten(input, 10)
		if result != input {
			t.Errorf("Shorten(%q, 10) = %q, expected %q", input, result, input)
		}
	})

	t.Run("string_equal_to_max", func(t *testing.T) {
		input := "hello"
		result := Shorten(input, 5)
		if result != input {
			t.Errorf("Shorten(%q, 5) = %q, expected %q", input, result, input)
		}
	})

	t.Run("string_longer_than_max", func(t *testing.T) {
		input := "hello world"
		result := Shorten(input, 5)
		expected := "hello"
		if result != expected {
			t.Errorf("Shorten(%q, 5) = %q, expected %q", input, result, expected)
		}
	})

	t.Run("unicode_shorten", func(t *testing.T) {
		input := "你好世界测试"
		result := Shorten(input, 9) // Each Chinese character is 3 bytes in UTF-8
		expected := "你好世"
		if result != expected {
			t.Errorf("Shorten(%q, 9) = %q, expected %q", input, result, expected)
		}
	})
}

func TestShortenShow(t *testing.T) {
	t.Run("negative_max", func(t *testing.T) {
		result := ShortenShow("hello", -1)
		expected := "..."
		if result != expected {
			t.Errorf("ShortenShow(\"hello\", -1) = %q, expected %q", result, expected)
		}
	})

	t.Run("max_less_than_3", func(t *testing.T) {
		result := ShortenShow("hello", 2)
		expected := "..."
		if result != expected {
			t.Errorf("ShortenShow(\"hello\", 2) = %q, expected %q", result, expected)
		}
	})

	t.Run("string_shorter_than_max", func(t *testing.T) {
		input := "hello"
		result := ShortenShow(input, 10)
		if result != input {
			t.Errorf("ShortenShow(%q, 10) = %q, expected %q", input, result, input)
		}
	})

	t.Run("string_equal_to_max", func(t *testing.T) {
		input := "hello"
		result := ShortenShow(input, 5)
		if result != input {
			t.Errorf("ShortenShow(%q, 5) = %q, expected %q", input, result, input)
		}
	})

	t.Run("string_longer_than_max", func(t *testing.T) {
		input := "hello world"
		result := ShortenShow(input, 8)
		expected := "hello..."
		if result != expected {
			t.Errorf("ShortenShow(%q, 8) = %q, expected %q", input, result, expected)
		}
	})

	t.Run("max_exactly_3", func(t *testing.T) {
		input := "hello"
		result := ShortenShow(input, 3)
		expected := "..."
		if result != expected {
			t.Errorf("ShortenShow(%q, 3) = %q, expected %q", input, result, expected)
		}
	})
}

func TestIsUpper(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected bool
	}{
		{"HELLO", true},
		{"hello", false},
		{"Hello", false},
		{"HELLO123", true},
		{"", true},
		{"123", true},
		{"!@#$%", true},
		{[]rune("HELLO"), true},
		{[]rune("hello"), false},
		{[]rune("Hello"), false},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			var result bool
			switch v := tc.input.(type) {
			case string:
				result = IsUpper(v)
			case []rune:
				result = IsUpper(v)
			}
			if result != tc.expected {
				t.Errorf("IsUpper(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestIsDigit(t *testing.T) {
	testCases := []struct {
		input    interface{}
		expected bool
	}{
		{"123", true},
		{"hello", false},
		{"123abc", false},
		{"", true},
		{"0", true},
		{"９８７", true}, // Full-width digits
		{[]rune("123"), true},
		{[]rune("hello"), false},
		{[]rune(""), true},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			var result bool
			switch v := tc.input.(type) {
			case string:
				result = IsDigit(v)
			case []rune:
				result = IsDigit(v)
			}
			if result != tc.expected {
				t.Errorf("IsDigit(%v) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"a", "a"},
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"12345", "54321"},
		{"你好世界", "界世好你"},
		{"Hello World", "dlroW olleH"},
		{"ab😀cd", "dc😀ba"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := Reverse(tc.input)
			if result != tc.expected {
				t.Errorf("Reverse(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestQuote(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", `"hello"`},
		{"", `""`},
		{"hello\nworld", `"hello\nworld"`},
		{"hello\"world", `"hello\"world"`},
		{"hello\\world", `"hello\\world"`},
		{"你好", `"你好"`},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := Quote(tc.input)
			if result != tc.expected {
				t.Errorf("Quote(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestQuotePure(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "hello"},
		{"", ""},
		{"hello\nworld", "hello\\nworld"},
		{"hello\"world", "hello\\\"world"},
		{"hello\\world", "hello\\\\world"},
		{"你好", "你好"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := QuotePure(tc.input)
			if result != tc.expected {
				t.Errorf("QuotePure(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

// Benchmark tests
func BenchmarkToString(b *testing.B) {
	data := []byte("hello world benchmark test string")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToString(data)
	}
}

func BenchmarkToBytes(b *testing.B) {
	data := "hello world benchmark test string"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToBytes(data)
	}
}

func BenchmarkCamel2Snake(b *testing.B) {
	data := "HTTPSConnectionPoolMaxSize"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Camel2Snake(data)
	}
}

func BenchmarkSnake2Camel(b *testing.B) {
	data := "http_connection_pool_max_size"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Snake2Camel(data)
	}
}

func BenchmarkReverse(b *testing.B) {
	data := "hello world benchmark test string with unicode 你好世界"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reverse(data)
	}
}

// Edge case tests
func TestStringEdgeCases(t *testing.T) {
	t.Run("very_long_string", func(t *testing.T) {
		longStr := strings.Repeat("a", 10000)
		result := Reverse(longStr)
		if len(result) != len(longStr) {
			t.Errorf("Reverse long string length mismatch: got %d, expected %d", len(result), len(longStr))
		}
	})

	t.Run("unicode_edge_cases", func(t *testing.T) {
		// Test with various Unicode categories
		testStrings := []string{
			"🙂😀🎉", // Emojis
			"Ñoël",    // Accented characters  
			"ٱلْعَرَبِيَّة", // Arabic
			"русский", // Cyrillic
			"中文",     // Chinese
			"🇺🇸🇨🇳",   // Flag emojis
		}

		for _, s := range testStrings {
			// Test various functions with Unicode
			_ = Reverse(s)
			_ = ToSnake(s)
			_ = ToCamel(s)
			_ = IsUpper(s)
			_ = IsDigit(s)
		}
	})
}