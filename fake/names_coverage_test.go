package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试FirstName函数的各种情况
func TestFirstNameCoverage(t *testing.T) {
	// 测试不同性别的情况
	genders := []Gender{
		GenderMale,
		GenderFemale,
		"", // 空性别，应该随机选择
	}

	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageFrench,
		LanguageRussian,
		LanguageSpanish,
		LanguagePortuguese,
	}

	for _, gender := range genders {
		for _, lang := range languages {
			t.Run("gender_"+string(gender)+"_lang_"+string(lang), func(t *testing.T) {
				f := New(WithGender(gender), WithLanguage(lang))
				firstName := f.FirstName()
				assert.NotEmpty(t, firstName)
			})
		}
	}

	// 测试全局便捷函数
	t.Run("global_FirstName", func(t *testing.T) {
		firstName := FirstName()
		assert.NotEmpty(t, firstName)
	})

	// 测试批量生成
	t.Run("BatchFirstNames", func(t *testing.T) {
		f := New()
		firstNames := f.BatchFirstNames(10)
		assert.Len(t, firstNames, 10)
		for _, name := range firstNames {
			assert.NotEmpty(t, name)
		}
	})
}

// 测试LastName函数的各种情况
func TestLastNameCoverage(t *testing.T) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageFrench,
		LanguageRussian,
		LanguageSpanish,
		LanguagePortuguese,
	}

	for _, lang := range languages {
		t.Run("lang_"+string(lang), func(t *testing.T) {
			f := New(WithLanguage(lang))
			lastName := f.LastName()
			assert.NotEmpty(t, lastName)
		})
	}

	// 测试全局便捷函数
	t.Run("global_LastName", func(t *testing.T) {
		lastName := LastName()
		assert.NotEmpty(t, lastName)
	})

	// 测试批量生成
	t.Run("BatchLastNames", func(t *testing.T) {
		f := New()
		lastNames := f.BatchLastNames(10)
		assert.Len(t, lastNames, 10)
		for _, name := range lastNames {
			assert.NotEmpty(t, name)
		}
	})
}

// 测试FullName函数的各种情况
func TestFullNameCoverage(t *testing.T) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageFrench,
	}

	for _, lang := range languages {
		t.Run("lang_"+string(lang), func(t *testing.T) {
			f := New(WithLanguage(lang))
			fullName := f.FullName()
			assert.NotEmpty(t, fullName)
		})
	}

	// 测试全局便捷函数
	t.Run("global_FullName", func(t *testing.T) {
		fullName := FullName()
		assert.NotEmpty(t, fullName)
	})
}

// 测试NamePrefix函数的各种情况
func TestNamePrefixCoverage(t *testing.T) {
	genders := []Gender{
		GenderMale,
		GenderFemale,
		"", // 空性别，应该随机选择
	}

	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
		LanguageFrench,
		LanguageRussian, // 应该返回空字符串，因为不支持
		"invalid_language", // 无效语言，应该返回空字符串
	}

	for _, gender := range genders {
		for _, lang := range languages {
			t.Run("gender_"+string(gender)+"_lang_"+string(lang), func(t *testing.T) {
				f := New(WithGender(gender), WithLanguage(lang))
				prefix := f.NamePrefix()
				// 对于不支持的语言，prefix应该是空字符串
				if lang == LanguageRussian || lang == "invalid_language" {
					assert.Empty(t, prefix)
				} else {
					assert.NotEmpty(t, prefix)
				}
			})
		}
	}

	// 测试全局便捷函数
	t.Run("global_NamePrefix", func(t *testing.T) {
		prefix := NamePrefix()
		assert.NotEmpty(t, prefix)
	})
}

// 测试NameSuffix函数的各种情况
func TestNameSuffixCoverage(t *testing.T) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageFrench,
	}

	for _, lang := range languages {
		t.Run("lang_"+string(lang), func(t *testing.T) {
			f := New(WithLanguage(lang))
			// 多次调用，增加触发后缀的概率
			for i := 0; i < 20; i++ {
				suffix := f.NameSuffix()
				// 对于非英语语言，suffix应该是空字符串
				if lang != LanguageEnglish {
					assert.Empty(t, suffix)
				}
			}
		})
	}

	// 测试全局便捷函数
	t.Run("global_NameSuffix", func(t *testing.T) {
		suffix := NameSuffix()
		// 可能为空，所以不做断言
		_ = suffix
	})
}

// 测试FormattedName函数的各种情况
func TestFormattedNameCoverage(t *testing.T) {
	languages := []Language{
		LanguageEnglish,
		LanguageChineseSimplified,
		LanguageChineseTraditional,
	}

	for _, lang := range languages {
		t.Run("lang_"+string(lang), func(t *testing.T) {
			f := New(WithLanguage(lang))
			formattedName := f.FormattedName()
			assert.NotEmpty(t, formattedName)
		})
	}

	// 测试全局便捷函数
	t.Run("global_FormattedName", func(t *testing.T) {
		formattedName := FormattedName()
		assert.NotEmpty(t, formattedName)
	})
}
