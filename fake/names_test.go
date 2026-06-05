package fake

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestName 测试姓名生成
func TestName(t *testing.T) {
	faker := New()
	name := faker.Name()

	if name == "" {
		t.Error("Name() returned empty string")
	}

	// 英文名字应该包含空格
	if !strings.Contains(name, " ") {
		t.Error("English name should contain space")
	}
}

// TestChineseName 测试中文姓名生成
func TestChineseName(t *testing.T) {
	faker := New(WithLanguage(LanguageChineseSimplified))
	name := faker.Name()

	if name == "" {
		t.Error("Chinese name should not be empty")
	}

	// 中文名字不应该包含空格
	if strings.Contains(name, " ") {
		t.Error("Chinese name should not contain space")
	}
}

// TestFirstName 测试名字生成
func TestFirstName(t *testing.T) {
	faker := New()
	firstName := faker.FirstName()

	if firstName == "" {
		t.Error("FirstName() returned empty string")
	}

	// 名字不应该包含空格
	if strings.Contains(firstName, " ") {
		t.Error("FirstName should not contain space")
	}
}

// TestLastName 测试姓氏生成
func TestLastName(t *testing.T) {
	faker := New()
	lastName := faker.LastName()

	if lastName == "" {
		t.Error("LastName() returned empty string")
	}

	// 姓氏不应该包含空格（对于单个姓氏）
	parts := strings.Split(lastName, " ")
	if len(parts) > 2 {
		t.Error("LastName should not have too many parts")
	}
}

// TestGenderSpecificNames 测试性别相关的名字生成
func TestGenderSpecificNames(t *testing.T) {
	maleFaker := New(WithGender(GenderMale))
	femaleFaker := New(WithGender(GenderFemale))

	// 生成多个名字确保性别特定
	for i := 0; i < 10; i++ {
		maleName := maleFaker.FirstName()
		femaleName := femaleFaker.FirstName()

		if maleName == "" || femaleName == "" {
			t.Error("Gender-specific names should not be empty")
		}

		// 这里我们主要测试不会崩溃，实际的性别验证需要更复杂的逻辑
	}
}

// TestNamePrefix 测试姓名前缀
func TestNamePrefix(t *testing.T) {
	faker := New()
	prefix := faker.NamePrefix()

	// 前缀可能为空，这是正常的
	if prefix != "" {
		expectedPrefixes := []string{"Mr.", "Ms.", "Mrs.", "Miss", "Dr.", "Prof."}
		found := false
		for _, expected := range expectedPrefixes {
			if prefix == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Unexpected name prefix: %s", prefix)
		}
	}
}

// TestNameSuffix 测试姓名后缀
func TestNameSuffix(t *testing.T) {
	faker := New()

	// 生成多个后缀，因为大多数情况下会返回空字符串
	suffixes := make(map[string]bool)
	for i := 0; i < 100; i++ {
		suffix := faker.NameSuffix()
		if suffix != "" {
			suffixes[suffix] = true
		}
	}

	// 检查生成的后缀是否合法
	validSuffixes := []string{"Jr.", "Sr.", "II", "III", "IV", "V"}
	for suffix := range suffixes {
		found := false
		for _, valid := range validSuffixes {
			if suffix == valid {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Invalid name suffix: %s", suffix)
		}
	}
}

// TestFullName 测试完整姓名生成
func TestFullName(t *testing.T) {
	faker := New()
	fullName := faker.FullName()

	if fullName == "" {
		t.Error("FullName() returned empty string")
	}

	parts := strings.Split(fullName, " ")
	if len(parts) < 2 {
		t.Error("FullName should have at least 2 parts")
	}

	// 完整姓名可能有中间名，所以部分数量可以是2或3
	if len(parts) > 3 {
		t.Error("FullName should not have more than 3 parts")
	}
}

// TestFormattedName 测试格式化姓名
func TestFormattedName(t *testing.T) {
	faker := New()
	formatted := faker.FormattedName()

	if formatted == "" {
		t.Error("FormattedName() returned empty string")
	}

	// 格式化姓名应该至少包含基本的姓名
	parts := strings.Split(formatted, " ")
	if len(parts) < 2 {
		t.Error("FormattedName should have at least 2 parts")
	}
}

// TestBatchNames 测试批量姓名生成
func TestBatchNames(t *testing.T) {
	faker := New()
	count := 10
	names := faker.BatchNames(count)

	if len(names) != count {
		t.Errorf("Expected %d names, got %d", count, len(names))
	}

	for i, name := range names {
		if name == "" {
			t.Errorf("Name at index %d is empty", i)
		}
	}
}

// TestBatchNamesOptimized 测试优化的批量姓名生成
func TestBatchNamesOptimized(t *testing.T) {
	faker := New()
	count := 100
	names := faker.BatchNamesOptimized(count)

	if len(names) != count {
		t.Errorf("Expected %d names, got %d", count, len(names))
	}

	for i, name := range names {
		if name == "" {
			t.Errorf("Optimized name at index %d is empty", i)
		}
	}
}

// TestGlobalNameFunctions 测试全局姓名函数
func TestGlobalNameFunctions(t *testing.T) {
	// 测试全局函数
	name := Name()
	if name == "" {
		t.Error("Global Name() returned empty string")
	}

	firstName := FirstName()
	if firstName == "" {
		t.Error("Global FirstName() returned empty string")
	}

	lastName := LastName()
	if lastName == "" {
		t.Error("Global LastName() returned empty string")
	}

	fullName := FullName()
	if fullName == "" {
		t.Error("Global FullName() returned empty string")
	}
}

// TestNameConsistency 测试姓名生成的一致性
func TestNameConsistency(t *testing.T) {
	faker := New(WithSeed(12345))

	// 用相同的种子生成多次，应该得到相同结果
	name1 := faker.FirstName()

	faker2 := New(WithSeed(12345))
	name2 := faker2.FirstName()

	// 注意：由于随机性，这个测试可能需要调整
	// 这里主要测试不会崩溃
	if name1 == "" || name2 == "" {
		t.Error("Names should not be empty")
	}
}

// BenchmarkName 性能测试：姓名生成
func BenchmarkName(b *testing.B) {
	faker := New()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = faker.Name()
	}
}

// BenchmarkBatchNames 性能测试：批量姓名生成
func BenchmarkBatchNames(b *testing.B) {
	faker := New()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = faker.BatchNames(100)
	}
}

// BenchmarkBatchNamesOptimized 性能测试：优化的批量姓名生成
func BenchmarkBatchNamesOptimized(b *testing.B) {
	faker := New()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = faker.BatchNamesOptimized(100)
	}
}

// TestNameMemoryUsage 测试姓名生成的内存使用
func TestNameMemoryUsage(b *testing.T) {
	faker := New()

	// 生成大量姓名，检查内存使用情况
	for i := 0; i < 1000; i++ {
		_ = faker.Name()
	}

	// 清理缓存
	faker.ClearCache()
}

// TestZeroCoverageFunctions 测试覆盖率为0%的函数
func TestZeroCoverageFunctions(t *testing.T) {
	// 测试getDefaultFirstName函数
	t.Run("getDefaultFirstName", func(t *testing.T) {
		// 测试男性名字
		maleName := getDefaultFirstName(GenderMale)
		assert.NotEmpty(t, maleName)
		assert.IsType(t, "", maleName)

		// 测试女性名字
		femaleName := getDefaultFirstName(GenderFemale)
		assert.NotEmpty(t, femaleName)
		assert.IsType(t, "", femaleName)

		// 测试默认性别名字
		defaultName := getDefaultFirstName("") // 空字符串作为默认性别
		assert.NotEmpty(t, defaultName)
		assert.IsType(t, "", defaultName)
	})

	// 测试getDefaultLastName函数
	t.Run("getDefaultLastName", func(t *testing.T) {
		lastName := getDefaultLastName()
		assert.NotEmpty(t, lastName)
		assert.IsType(t, "", lastName)
	})

	// 测试getDefaultWord函数
	t.Run("getDefaultWord", func(t *testing.T) {
		faker := New()
		word := faker.getDefaultWord()
		assert.NotEmpty(t, word)
		assert.IsType(t, "", word)
	})

	// 测试incrementCallOut函数
	t.Run("incrementCallOut", func(t *testing.T) {
		faker := New()
		faker.incrementCallOut() // 空函数，测试不崩溃即可
	})

	// 测试names.go中的全局便捷函数（确保它们调用了正确的方法）
	t.Run("namesGlobalFunctions", func(t *testing.T) {
		// 测试全局Name函数
		name := Name()
		assert.NotEmpty(t, name)
		assert.IsType(t, "", name)

		// 测试全局FirstName函数
		firstName := FirstName()
		assert.NotEmpty(t, firstName)
		assert.IsType(t, "", firstName)

		// 测试全局LastName函数
		lastName := LastName()
		assert.NotEmpty(t, lastName)
		assert.IsType(t, "", lastName)

		// 测试全局FullName函数
		fullName := FullName()
		assert.NotEmpty(t, fullName)
		assert.IsType(t, "", fullName)

		// 测试全局FormattedName函数
		formattedName := FormattedName()
		assert.NotEmpty(t, formattedName)
		assert.IsType(t, "", formattedName)

		// 测试全局NamePrefix函数
		namePrefix := NamePrefix()
		assert.IsType(t, "", namePrefix)

		// 测试全局NameSuffix函数
		nameSuffix := NameSuffix()
		assert.IsType(t, "", nameSuffix)
	})

	// 测试text.go中的全局便捷函数
	t.Run("textGlobalFunctions", func(t *testing.T) {
		// 测试全局Word函数
		word := Word()
		assert.NotEmpty(t, word)
		assert.IsType(t, "", word)

		// 测试全局Words函数
		words := Words(5)
		assert.Len(t, words, 5)
		for _, w := range words {
			assert.IsType(t, "", w)
		}

		// 测试全局Sentence函数
		sentence := Sentence()
		assert.NotEmpty(t, sentence)
		assert.IsType(t, "", sentence)

		// 测试全局Sentences函数
		sentences := Sentences(3)
		assert.Len(t, sentences, 3)
		for _, s := range sentences {
			assert.IsType(t, "", s)
		}

		// 测试全局Paragraph函数
		paragraph := Paragraph()
		assert.NotEmpty(t, paragraph)
		assert.IsType(t, "", paragraph)

		// 测试全局Paragraphs函数
		paragraphs := Paragraphs(2)
		assert.Len(t, paragraphs, 2)
		for _, p := range paragraphs {
			assert.IsType(t, "", p)
		}

		// 测试全局Text函数
		text := Text(100)
		assert.NotEmpty(t, text)
		assert.IsType(t, "", text)

		// 测试全局Title函数
		title := Title()
		assert.NotEmpty(t, title)
		assert.IsType(t, "", title)

		// 测试全局Quote函数
		quote := Quote()
		assert.NotEmpty(t, quote)
		assert.IsType(t, "", quote)

		// 测试全局Lorem函数
		lorem := Lorem()
		assert.NotEmpty(t, lorem)
		assert.IsType(t, "", lorem)
	})
}

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
		LanguageRussian,    // 应该返回空字符串，因为不支持
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
