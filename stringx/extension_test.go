package stringx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimSpaceAll(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", TrimSpaceAll(""))
	})

	t.Run("no_spaces", func(t *testing.T) {
		assert.Equal(t, "hello", TrimSpaceAll("hello"))
	})

	t.Run("spaces_only", func(t *testing.T) {
		assert.Equal(t, "", TrimSpaceAll("     "))
	})

	t.Run("internal_spaces", func(t *testing.T) {
		assert.Equal(t, "helloworld", TrimSpaceAll("hello world"))
	})

	t.Run("multiple_spaces", func(t *testing.T) {
		assert.Equal(t, "helloworld", TrimSpaceAll("hello   world"))
	})

	t.Run("tabs_and_newlines", func(t *testing.T) {
		assert.Equal(t, "helloworld", TrimSpaceAll("hello\t\nworld"))
	})

	t.Run("mixed_whitespace", func(t *testing.T) {
		assert.Equal(t, "abc", TrimSpaceAll(" a b c "))
	})

	t.Run("unicode_spaces", func(t *testing.T) {
		// 测试 Unicode 空白字符
		result := TrimSpaceAll("hello​world")
		assert.Contains(t, result, "hello")
		assert.Contains(t, result, "world")
	})
}

func TestNormalize(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", Normalize("", 0))
		assert.Equal(t, "", Normalize("", 1))
		assert.Equal(t, "", Normalize("", 2))
		assert.Equal(t, "", Normalize("", 3))
	})

	t.Run("nfc", func(t *testing.T) {
		// NFC 组合字符
		input := "é" // 可能是组合字符或预组合字符
		result := Normalize(input, 0)
		assert.NotEmpty(t, result)
	})

	t.Run("nfd", func(t *testing.T) {
		// NFD 分解字符
		input := "é"
		result := Normalize(input, 1)
		assert.NotEmpty(t, result)
	})
}

func TestBase64Encode(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", Base64Encode(""))
	})

	t.Run("basic_encoding", func(t *testing.T) {
		assert.Equal(t, "SGVsbG8gV29ybGQ=", Base64Encode("Hello World"))
	})

	t.Run("unicode", func(t *testing.T) {
		result := Base64Encode("你好")
		assert.NotEmpty(t, result)
	})
}

func TestBase64Decode(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		result, err := Base64Decode("")
		assert.NoError(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("basic_decoding", func(t *testing.T) {
		result, err := Base64Decode("SGVsbG8gV29ybGQ=")
		assert.NoError(t, err)
		assert.Equal(t, "Hello World", result)
	})

	t.Run("invalid_input", func(t *testing.T) {
		_, err := Base64Decode("invalid@base64!")
		assert.Error(t, err)
	})
}

func TestBase64URLEncode(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", Base64URLEncode(""))
	})

	t.Run("basic_encoding", func(t *testing.T) {
		result := Base64URLEncode("Hello")
		assert.NotEmpty(t, result)
		// URL safe encoding should not have + or /
		assert.NotContains(t, result, "+")
		assert.NotContains(t, result, "/")
	})
}

func TestBase64URLDecode(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		result, err := Base64URLDecode("")
		assert.NoError(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("basic_decoding", func(t *testing.T) {
		encoded := Base64URLEncode("Hello World")
		result, err := Base64URLDecode(encoded)
		assert.NoError(t, err)
		assert.Equal(t, "Hello World", result)
	})
}

func TestHexEncode(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", HexEncode(""))
	})

	t.Run("basic_encoding", func(t *testing.T) {
		assert.Equal(t, "48656c6c6f", HexEncode("Hello"))
	})

	t.Run("chinese", func(t *testing.T) {
		result := HexEncode("你好")
		assert.NotEmpty(t, result)
		assert.Equal(t, len(result)%2, 0)
	})
}

func TestHexDecode(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		result, err := HexDecode("")
		assert.NoError(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("basic_decoding", func(t *testing.T) {
		result, err := HexDecode("48656c6c6f")
		assert.NoError(t, err)
		assert.Equal(t, "Hello", result)
	})

	t.Run("invalid_hex", func(t *testing.T) {
		_, err := HexDecode("xyz")
		assert.Error(t, err)
	})
}

func TestMask(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", Mask("", 2))
	})

	t.Run("negative_visible", func(t *testing.T) {
		assert.Equal(t, "****", Mask("abcd", -1))
	})

	t.Run("visible_zero", func(t *testing.T) {
		assert.Equal(t, "****", Mask("abcd", 0))
	})

	t.Run("visible_greater_than_half", func(t *testing.T) {
		assert.Equal(t, "abcd", Mask("abcd", 3))
	})

	t.Run("basic_masking", func(t *testing.T) {
		assert.Equal(t, "12****78", Mask("12345678", 2))
	})

	t.Run("unicode_masking", func(t *testing.T) {
		result := Mask("你好世界", 1)
		assert.Contains(t, result, "你")
		assert.Contains(t, result, "*")
		assert.Contains(t, result, "界")
	})

	t.Run("exact_half", func(t *testing.T) {
		result := Mask("123456", 2)
		assert.Contains(t, result, "12")
		assert.Contains(t, result, "56")
		assert.Contains(t, result, "*")
	})
}

func TestMaskEmail(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", MaskEmail(""))
	})

	t.Run("valid_email", func(t *testing.T) {
		assert.Equal(t, "u***@example.com", MaskEmail("user@example.com"))
	})

	t.Run("short_username", func(t *testing.T) {
		assert.Equal(t, "a***@example.com", MaskEmail("a@example.com"))
	})

	t.Run("invalid_format", func(t *testing.T) {
		// 无 @ 符号，使用通用掩码
		result := MaskEmail("noatsign")
		assert.Contains(t, result, "*")
	})
}

func TestMaskPhone(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", MaskPhone(""))
	})

	t.Run("normal_phone", func(t *testing.T) {
		assert.Equal(t, "138****5678", MaskPhone("13812345678"))
	})

	t.Run("short_phone", func(t *testing.T) {
		result := MaskPhone("1234567")
		assert.NotEmpty(t, result)
		assert.Contains(t, result, "*")
	})

	t.Run("long_phone", func(t *testing.T) {
		result := MaskPhone("1234567890")
		assert.Contains(t, result, "123")
		assert.Contains(t, result, "7890")
		assert.Contains(t, result, "*")
	})
}

func TestEditDistance(t *testing.T) {
	t.Run("empty_strings", func(t *testing.T) {
		assert.Equal(t, 0, EditDistance("", ""))
	})

	t.Run("identical", func(t *testing.T) {
		assert.Equal(t, 0, EditDistance("hello", "hello"))
	})

	t.Run("completely_different", func(t *testing.T) {
		assert.Equal(t, 4, EditDistance("hello", "world"))
	})

	t.Run("one_insertion", func(t *testing.T) {
		assert.Equal(t, 1, EditDistance("hell", "hello"))
	})

	t.Run("one_deletion", func(t *testing.T) {
		assert.Equal(t, 1, EditDistance("hello", "hell"))
	})

	t.Run("one_substitution", func(t *testing.T) {
		assert.Equal(t, 1, EditDistance("hello", "hallo"))
	})

	t.Run("empty_to_nonempty", func(t *testing.T) {
		assert.Equal(t, 5, EditDistance("", "hello"))
		assert.Equal(t, 5, EditDistance("hello", ""))
	})

	t.Run("unicode", func(t *testing.T) {
		distance := EditDistance("你好", "您好")
		assert.Greater(t, distance, 0)
	})
}

func TestSimilarity(t *testing.T) {
	t.Run("identical", func(t *testing.T) {
		assert.Equal(t, 1.0, Similarity("hello", "hello"))
	})

	t.Run("empty_strings", func(t *testing.T) {
		assert.Equal(t, 1.0, Similarity("", ""))
	})

	t.Run("completely_different", func(t *testing.T) {
		similarity := Similarity("abc", "xyz")
		assert.GreaterOrEqual(t, similarity, 0.0)
		assert.LessOrEqual(t, similarity, 1.0)
	})

	t.Run("three_edits", func(t *testing.T) {
		assert.Equal(t, 3, EditDistance("abc", "xyz"))
	})

	t.Run("one_edit_difference", func(t *testing.T) {
		similarity := Similarity("hello", "hallo")
		assert.InDelta(t, 0.8, similarity, 0.01)
	})
}

func TestSlugify(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", Slugify(""))
	})

	t.Run("basic", func(t *testing.T) {
		assert.Equal(t, "hello-world", Slugify("Hello World"))
	})

	t.Run("with_exclamation", func(t *testing.T) {
		assert.Equal(t, "hello-world", Slugify("Hello World!"))
	})

	t.Run("multiple_spaces", func(t *testing.T) {
		assert.Equal(t, "hello-world", Slugify("Hello   World"))
	})

	t.Run("with_special_chars", func(t *testing.T) {
		assert.Equal(t, "test-string", Slugify("Test @#$% String"))
	})

	t.Run("underscores_to_hyphens", func(t *testing.T) {
		assert.Equal(t, "hello-world-test", Slugify("hello_world_test"))
	})

	t.Run("mixed_case", func(t *testing.T) {
		assert.Equal(t, "hello-world", Slugify("HeLLo WoRLd"))
	})

	t.Run("numbers", func(t *testing.T) {
		assert.Equal(t, "test-123", Slugify("Test 123"))
	})
}

func TestRemoveHyphens(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", RemoveHyphens(""))
	})

	t.Run("no_hyphens", func(t *testing.T) {
		assert.Equal(t, "hello", RemoveHyphens("hello"))
	})

	t.Run("with_hyphens", func(t *testing.T) {
		assert.Equal(t, "helloworld", RemoveHyphens("hello-world"))
	})

	t.Run("multiple_hyphens", func(t *testing.T) {
		assert.Equal(t, "helloworldtest", RemoveHyphens("hello-world-test"))
	})

	t.Run("hyphens_at_edges", func(t *testing.T) {
		assert.Equal(t, "helloworld", RemoveHyphens("-hello-world-"))
	})
}

func TestNormalizeHyphens(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", NormalizeHyphens(""))
	})

	t.Run("standard_hyphen", func(t *testing.T) {
		assert.Equal(t, "hello-world", NormalizeHyphens("hello-world"))
	})

	t.Run("en_dash", func(t *testing.T) {
		result := NormalizeHyphens("hello–world") // en dash
		assert.Equal(t, "hello-world", result)
	})

	t.Run("em_dash", func(t *testing.T) {
		result := NormalizeHyphens("hello—world") // em dash
		assert.Equal(t, "hello-world", result)
	})

	t.Run("mixed_dashes", func(t *testing.T) {
		result := NormalizeHyphens("hello–world—test")
		assert.Equal(t, "hello-world-test", result)
	})
}

func TestCountWords(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, 0, CountWords(""))
	})

	t.Run("single_word", func(t *testing.T) {
		assert.Equal(t, 1, CountWords("hello"))
	})

	t.Run("multiple_words", func(t *testing.T) {
		assert.Equal(t, 3, CountWords("hello world test"))
	})

	t.Run("multiple_spaces", func(t *testing.T) {
		assert.Equal(t, 3, CountWords("hello   world   test"))
	})

	t.Run("leading_trailing_spaces", func(t *testing.T) {
		assert.Equal(t, 2, CountWords("  hello world  "))
	})

	t.Run("tabs_and_newlines", func(t *testing.T) {
		assert.Equal(t, 3, CountWords("hello\tworld\ntest"))
	})

	t.Run("chinese", func(t *testing.T) {
		// 中文分词基于空格
		assert.Equal(t, 1, CountWords("你好世界"))
	})
}

func TestCountLines(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, 0, CountLines(""))
	})

	t.Run("single_line", func(t *testing.T) {
		assert.Equal(t, 1, CountLines("hello"))
	})

	t.Run("multiple_lines", func(t *testing.T) {
		assert.Equal(t, 3, CountLines("hello\nworld\ntest"))
	})

	t.Run("with_trailing_newline", func(t *testing.T) {
		assert.Equal(t, 2, CountLines("hello\n"))
	})

	t.Run("unix_vs_windows", func(t *testing.T) {
		assert.Equal(t, 2, CountLines("hello\r\nworld"))
	})
}

func TestCountRunes(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, 0, CountRunes(""))
	})

	t.Run("ascii", func(t *testing.T) {
		assert.Equal(t, 5, CountRunes("hello"))
	})

	t.Run("unicode", func(t *testing.T) {
		assert.Equal(t, 4, CountRunes("你好世界"))
	})

	t.Run("emoji", func(t *testing.T) {
		count := CountRunes("hello😀world")
		assert.GreaterOrEqual(t, count, 5)
	})
}

func TestToLower(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", ToLower(""))
	})

	t.Run("already_lower", func(t *testing.T) {
		assert.Equal(t, "hello", ToLower("hello"))
	})

	t.Run("mixed_case", func(t *testing.T) {
		assert.Equal(t, "hello", ToLower("HeLLo"))
	})

	t.Run("unicode", func(t *testing.T) {
		result := ToLower("HELLO你好")
		assert.Contains(t, result, "hello")
		assert.Contains(t, result, "你好")
	})
}

func TestToUpper(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", ToUpper(""))
	})

	t.Run("already_upper", func(t *testing.T) {
		assert.Equal(t, "HELLO", ToUpper("HELLO"))
	})

	t.Run("mixed_case", func(t *testing.T) {
		assert.Equal(t, "HELLO", ToUpper("HeLLo"))
	})

	t.Run("unicode", func(t *testing.T) {
		result := ToUpper("hello你好")
		assert.Contains(t, result, "HELLO")
		assert.Contains(t, result, "你好")
	})
}

func TestToTitleCase(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", ToTitleCase(""))
	})

	t.Run("single_word", func(t *testing.T) {
		assert.Equal(t, "Hello", ToTitleCase("hello"))
	})

	t.Run("multiple_words", func(t *testing.T) {
		result := ToTitleCase("hello world")
		assert.NotEqual(t, "hello world", result)
	})
}

func TestContains(t *testing.T) {
	t.Run("contains", func(t *testing.T) {
		assert.True(t, Contains("hello world", "wo"))
	})

	t.Run("not_contains", func(t *testing.T) {
		assert.False(t, Contains("hello world", "xyz"))
	})

	t.Run("empty_substring", func(t *testing.T) {
		assert.True(t, Contains("hello", ""))
	})
}

func TestHasPrefix(t *testing.T) {
	t.Run("has_prefix", func(t *testing.T) {
		assert.True(t, HasPrefix("hello world", "hello"))
	})

	t.Run("no_prefix", func(t *testing.T) {
		assert.False(t, HasPrefix("hello world", "world"))
	})

	t.Run("exact_match", func(t *testing.T) {
		assert.True(t, HasPrefix("hello", "hello"))
	})
}

func TestHasSuffix(t *testing.T) {
	t.Run("has_suffix", func(t *testing.T) {
		assert.True(t, HasSuffix("hello world", "world"))
	})

	t.Run("no_suffix", func(t *testing.T) {
		assert.False(t, HasSuffix("hello world", "hello"))
	})

	t.Run("exact_match", func(t *testing.T) {
		assert.True(t, HasSuffix("hello", "hello"))
	})
}

func TestTrimPrefix(t *testing.T) {
	t.Run("has_prefix", func(t *testing.T) {
		assert.Equal(t, " world", TrimPrefix("hello world", "hello"))
	})

	t.Run("no_prefix", func(t *testing.T) {
		assert.Equal(t, "hello world", TrimPrefix("hello world", "xyz"))
	})
}

func TestTrimSuffix(t *testing.T) {
	t.Run("has_suffix", func(t *testing.T) {
		assert.Equal(t, "hello ", TrimSuffix("hello world", "world"))
	})

	t.Run("no_suffix", func(t *testing.T) {
		assert.Equal(t, "hello world", TrimSuffix("hello world", "xyz"))
	})
}

func TestIndex(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		assert.Equal(t, 2, Index("hello world", "l"))
	})

	t.Run("not_found", func(t *testing.T) {
		assert.Equal(t, -1, Index("hello", "xyz"))
	})

	t.Run("empty_substring", func(t *testing.T) {
		assert.Equal(t, 0, Index("hello", ""))
	})
}

func TestLastIndex(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		assert.Equal(t, 9, LastIndex("hello world", "l"))
	})

	t.Run("not_found", func(t *testing.T) {
		assert.Equal(t, -1, LastIndex("hello", "xyz"))
	})
}

func TestSubstring(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", Substring("", 0, 5))
	})

	t.Run("positive_indices", func(t *testing.T) {
		assert.Equal(t, "ell", Substring("hello", 1, 4))
	})

	t.Run("negative_start", func(t *testing.T) {
		assert.Equal(t, "lo", Substring("hello", -2, 5))
	})

	t.Run("negative_end", func(t *testing.T) {
		assert.Equal(t, "ell", Substring("hello", 1, -1))
	})

	t.Run("both_negative", func(t *testing.T) {
		assert.Equal(t, "l", Substring("hello", -2, -1))
	})

	t.Run("out_of_bounds", func(t *testing.T) {
		assert.Equal(t, "hello", Substring("hello", 0, 100))
	})

	t.Run("invalid_range", func(t *testing.T) {
		assert.Equal(t, "", Substring("hello", 3, 1))
	})

	t.Run("unicode", func(t *testing.T) {
		assert.Equal(t, "好世", Substring("你好世界", 1, 3))
	})
}

func TestRemoveDuplicates(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", RemoveDuplicates(""))
	})

	t.Run("single_char", func(t *testing.T) {
		assert.Equal(t, "a", RemoveDuplicates("a"))
	})

	t.Run("no_duplicates", func(t *testing.T) {
		assert.Equal(t, "abc", RemoveDuplicates("abc"))
	})

	t.Run("with_duplicates", func(t *testing.T) {
		assert.Equal(t, "abca", RemoveDuplicates("aaabbbcccaaa"))
	})

	t.Run("all_same", func(t *testing.T) {
		assert.Equal(t, "a", RemoveDuplicates("aaaaa"))
	})
}

func TestReverseWords(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", ReverseWords(""))
	})

	t.Run("single_word", func(t *testing.T) {
		assert.Equal(t, "hello", ReverseWords("hello"))
	})

	t.Run("multiple_words", func(t *testing.T) {
		assert.Equal(t, "world hello", ReverseWords("hello world"))
	})

	t.Run("three_words", func(t *testing.T) {
		assert.Equal(t, "test world hello", ReverseWords("hello world test"))
	})

	t.Run("multiple_spaces", func(t *testing.T) {
		assert.Equal(t, "test world hello", ReverseWords("hello   world   test"))
	})
}

func TestCapitalize(t *testing.T) {
	t.Run("empty_string", func(t *testing.T) {
		assert.Equal(t, "", Capitalize(""))
	})

	t.Run("single_word", func(t *testing.T) {
		assert.Equal(t, "Hello", Capitalize("hello"))
		assert.Equal(t, "Hello", Capitalize("HELLO"))
	})

	t.Run("multiple_words", func(t *testing.T) {
		assert.Equal(t, "Hello world", Capitalize("hello WORLD"))
	})

	t.Run("single_char", func(t *testing.T) {
		assert.Equal(t, "A", Capitalize("a"))
		assert.Equal(t, "A", Capitalize("A"))
	})
}
