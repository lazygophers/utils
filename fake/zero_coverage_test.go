package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
