package fake

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBankAccount(t *testing.T) {
	faker := New()

	account := faker.BankAccount()
	assert.NotEmpty(t, account)
	assert.GreaterOrEqual(t, len(account), 8)
	assert.LessOrEqual(t, len(account), 17)
}

func TestIBAN(t *testing.T) {
	faker := New()

	iban := faker.IBAN()
	assert.NotEmpty(t, iban)
	assert.GreaterOrEqual(t, len(iban), 15)
	assert.LessOrEqual(t, len(iban), 34)
}

func TestSafeCreditCardNumber(t *testing.T) {
	faker := New()

	cc := faker.SafeCreditCardNumber()
	assert.NotEmpty(t, cc)
	assert.GreaterOrEqual(t, len(cc), 13)
	assert.LessOrEqual(t, len(cc), 19)
}

func TestBatchSSNs(t *testing.T) {
	faker := New()

	ssns := faker.BatchSSNs(5)
	assert.Len(t, ssns, 5)

	for _, ssn := range ssns {
		assert.NotEmpty(t, ssn)
	}
}

func TestBatchCreditCardNumbers(t *testing.T) {
	faker := New()

	ccs := faker.BatchCreditCardNumbers(5)
	assert.Len(t, ccs, 5)

	for _, cc := range ccs {
		assert.NotEmpty(t, cc)
	}
}

func TestBatchCreditCardInfos(t *testing.T) {
	faker := New()

	infos := faker.BatchCreditCardInfos(5)
	assert.Len(t, infos, 5)

	for _, info := range infos {
		assert.NotNil(t, info)
		assert.NotEmpty(t, info.Number)
		assert.NotEmpty(t, info.Type)
		assert.NotEmpty(t, info.ExpiryMonth)
		assert.NotEmpty(t, info.ExpiryYear)
		assert.NotEmpty(t, info.CVV)
	}
}

func TestBatchFirstNames(t *testing.T) {
	faker := New()

	names := faker.BatchFirstNames(5)
	assert.Len(t, names, 5)

	for _, name := range names {
		assert.NotEmpty(t, name)
	}
}

func TestBatchLastNames(t *testing.T) {
	faker := New()

	names := faker.BatchLastNames(5)
	assert.Len(t, names, 5)

	for _, name := range names {
		assert.NotEmpty(t, name)
	}
}

func TestBatchJobTitles(t *testing.T) {
	faker := New()

	titles := faker.BatchJobTitles(5)
	assert.Len(t, titles, 5)

	for _, title := range titles {
		assert.NotEmpty(t, title)
	}
}

func TestBatchCompanyInfos(t *testing.T) {
	faker := New()

	infos := faker.BatchCompanyInfos(5)
	assert.Len(t, infos, 5)

	for _, info := range infos {
		assert.NotNil(t, info)
		assert.NotEmpty(t, info.Name)
		assert.NotEmpty(t, info.Industry)
	}
}

func TestWord(t *testing.T) {
	faker := New()

	word := faker.Word()
	assert.NotEmpty(t, word)
}

func TestWords(t *testing.T) {
	faker := New()

	words := faker.Words(5)
	assert.Len(t, words, 5)

	for _, word := range words {
		assert.NotEmpty(t, word)
	}
}

func TestSentence(t *testing.T) {
	faker := New()

	sentence := faker.Sentence()
	assert.NotEmpty(t, sentence)
}

func TestSentences(t *testing.T) {
	faker := New()

	sentences := faker.Sentences(5)
	assert.Len(t, sentences, 5)

	for _, sentence := range sentences {
		assert.NotEmpty(t, sentence)
	}
}

func TestParagraph(t *testing.T) {
	faker := New()

	paragraph := faker.Paragraph()
	assert.NotEmpty(t, paragraph)
}

func TestParagraphs(t *testing.T) {
	faker := New()

	paragraphs := faker.Paragraphs(5)
	assert.Len(t, paragraphs, 5)

	for _, paragraph := range paragraphs {
		assert.NotEmpty(t, paragraph)
	}
}

func TestText(t *testing.T) {
	faker := New()

	text := faker.Text(100)
	assert.NotEmpty(t, text)
	assert.LessOrEqual(t, len(text), 100)
}

func TestTitle(t *testing.T) {
	faker := New()

	title := faker.Title()
	assert.NotEmpty(t, title)
}

func TestQuote(t *testing.T) {
	faker := New()

	quote := faker.Quote()
	assert.NotEmpty(t, quote)
}

func TestLorem(t *testing.T) {
	faker := New()

	lorem := faker.Lorem()
	assert.NotEmpty(t, lorem)
}

func TestLoremWords(t *testing.T) {
	faker := New()

	words := faker.LoremWords(10)
	assert.NotEmpty(t, words)
}

func TestLoremSentences(t *testing.T) {
	faker := New()

	sentences := faker.LoremSentences(3)
	assert.NotEmpty(t, sentences)
}

func TestLoremParagraphs(t *testing.T) {
	faker := New()

	paragraphs := faker.LoremParagraphs(2)
	assert.NotEmpty(t, paragraphs)
}

func TestArticle(t *testing.T) {
	faker := New()

	article := faker.Article()
	assert.NotEmpty(t, article)
}

func TestSlug(t *testing.T) {
	faker := New()

	slug := faker.Slug()
	assert.NotEmpty(t, slug)
}

func TestHashTag(t *testing.T) {
	faker := New()

	hashtag := faker.HashTag()
	assert.NotEmpty(t, hashtag)
	assert.Contains(t, hashtag, "#")
}

func TestHashTags(t *testing.T) {
	faker := New()

	hashtags := faker.HashTags(5)
	assert.Len(t, hashtags, 5)

	for _, tag := range hashtags {
		assert.NotEmpty(t, tag)
		assert.Contains(t, tag, "#")
	}
}

func TestTweet(t *testing.T) {
	faker := New()

	tweet := faker.Tweet()
	assert.NotEmpty(t, tweet)
}

func TestReview(t *testing.T) {
	faker := New()

	review := faker.Review()
	assert.NotEmpty(t, review)
}

func TestBatchWords(t *testing.T) {
	faker := New()

	words := faker.BatchWords(5)
	assert.Len(t, words, 5)

	for _, word := range words {
		assert.NotEmpty(t, word)
	}
}

func TestBatchSentences(t *testing.T) {
	faker := New()

	sentences := faker.BatchSentences(5)
	assert.Len(t, sentences, 5)

	for _, sentence := range sentences {
		assert.NotEmpty(t, sentence)
	}
}

func TestBatchParagraphs(t *testing.T) {
	faker := New()

	paragraphs := faker.BatchParagraphs(5)
	assert.Len(t, paragraphs, 5)

	for _, paragraph := range paragraphs {
		assert.NotEmpty(t, paragraph)
	}
}

func TestBatchTitles(t *testing.T) {
	faker := New()

	titles := faker.BatchTitles(5)
	assert.Len(t, titles, 5)

	for _, title := range titles {
		assert.NotEmpty(t, title)
	}
}

// TestGlobalTextFunctionsAdditional 测试text.go中的全局便捷函数
func TestGlobalTextFunctionsAdditional(t *testing.T) {
	// 测试Word全局函数
	word := Word()
	if word == "" {
		t.Error("Global Word() should not return empty string")
	}

	// 测试Words全局函数
	words := Words(5)
	if len(words) != 5 {
		t.Errorf("Global Words(5) should return 5 items, got %d", len(words))
	}

	// 测试Sentence全局函数
	sentence := Sentence()
	if sentence == "" {
		t.Error("Global Sentence() should not return empty string")
	}

	// 测试Sentences全局函数
	sentences := Sentences(3)
	if len(sentences) != 3 {
		t.Errorf("Global Sentences(3) should return 3 items, got %d", len(sentences))
	}

	// 测试Paragraph全局函数
	paragraph := Paragraph()
	if paragraph == "" {
		t.Error("Global Paragraph() should not return empty string")
	}

	// 测试Paragraphs全局函数
	paragraphs := Paragraphs(2)
	if len(paragraphs) != 2 {
		t.Errorf("Global Paragraphs(2) should return 2 items, got %d", len(paragraphs))
	}

	// 测试Text全局函数
	text := Text(100)
	if text == "" {
		t.Error("Global Text(100) should not return empty string")
	}

	// 测试Title全局函数
	title := Title()
	if title == "" {
		t.Error("Global Title() should not return empty string")
	}

	// 测试Quote全局函数
	quote := Quote()
	if quote == "" {
		t.Error("Global Quote() should not return empty string")
	}

	// 测试Lorem全局函数
	lorem := Lorem()
	if lorem == "" {
		t.Error("Global Lorem() should not return empty string")
	}

	// 测试LoremWords全局函数
	loremWords := LoremWords(10)
	if loremWords == "" {
		t.Error("Global LoremWords(10) should not return empty string")
	}

	// 测试LoremSentences全局函数
	loremSentences := LoremSentences(5)
	if loremSentences == "" {
		t.Error("Global LoremSentences(5) should not return empty string")
	}

	// 测试LoremParagraphs全局函数
	loremParagraphs := LoremParagraphs(2)
	if loremParagraphs == "" {
		t.Error("Global LoremParagraphs(2) should not return empty string")
	}

	// 测试Article全局函数
	article := Article()
	if article == "" {
		t.Error("Global Article() should not return empty string")
	}

	// 测试Slug全局函数
	slug := Slug()
	if slug == "" {
		t.Error("Global Slug() should not return empty string")
	}

	// 测试HashTag全局函数
	hashTag := HashTag()
	if hashTag == "" {
		t.Error("Global HashTag() should not return empty string")
	}

	// 测试HashTags全局函数
	hashTags := HashTags(5)
	if len(hashTags) != 5 {
		t.Errorf("Global HashTags(5) should return 5 items, got %d", len(hashTags))
	}

	// 测试Tweet全局函数
	tweet := Tweet()
	if tweet == "" {
		t.Error("Global Tweet() should not return empty string")
	}

	// 测试Review全局函数
	review := Review()
	if review == "" {
		t.Error("Global Review() should not return empty string")
	}
}

// TestTextMethodsAdditional 测试text.go中的实例方法
func TestTextMethodsAdditional(t *testing.T) {
	faker := New()

	// 测试Word方法
	word := faker.Word()
	if word == "" {
		t.Error("Word() should not return empty string")
	}

	// 测试Words方法
	words := faker.Words(5)
	if len(words) != 5 {
		t.Errorf("Words(5) should return 5 items, got %d", len(words))
	}

	// 测试Sentence方法
	sentence := faker.Sentence()
	if sentence == "" {
		t.Error("Sentence() should not return empty string")
	}

	// 测试Sentences方法
	sentences := faker.Sentences(3)
	if len(sentences) != 3 {
		t.Errorf("Sentences(3) should return 3 items, got %d", len(sentences))
	}

	// 测试Paragraph方法
	paragraph := faker.Paragraph()
	if paragraph == "" {
		t.Error("Paragraph() should not return empty string")
	}

	// 测试Paragraphs方法
	paragraphs := faker.Paragraphs(2)
	if len(paragraphs) != 2 {
		t.Errorf("Paragraphs(2) should return 2 items, got %d", len(paragraphs))
	}

	// 测试Text方法
	text := faker.Text(100)
	if text == "" {
		t.Error("Text(100) should not return empty string")
	}

	// 测试Title方法
	title := faker.Title()
	if title == "" {
		t.Error("Title() should not return empty string")
	}

	// 测试Quote方法
	quote := faker.Quote()
	if quote == "" {
		t.Error("Quote() should not return empty string")
	}

	// 测试Lorem方法
	lorem := faker.Lorem()
	if lorem == "" {
		t.Error("Lorem() should not return empty string")
	}

	// 测试LoremWords方法
	loremWords := faker.LoremWords(10)
	if loremWords == "" {
		t.Error("LoremWords(10) should not return empty string")
	}

	// 测试LoremSentences方法
	loremSentences := faker.LoremSentences(5)
	if loremSentences == "" {
		t.Error("LoremSentences(5) should not return empty string")
	}

	// 测试LoremParagraphs方法
	loremParagraphs := faker.LoremParagraphs(2)
	if loremParagraphs == "" {
		t.Error("LoremParagraphs(2) should not return empty string")
	}

	// 测试Article方法
	article := faker.Article()
	if article == "" {
		t.Error("Article() should not return empty string")
	}

	// 测试Slug方法
	slug := faker.Slug()
	if slug == "" {
		t.Error("Slug() should not return empty string")
	}

	// 测试HashTag方法
	hashTag := faker.HashTag()
	if hashTag == "" {
		t.Error("HashTag() should not return empty string")
	}

	// 测试HashTags方法
	hashTags := faker.HashTags(5)
	if len(hashTags) != 5 {
		t.Errorf("HashTags(5) should return 5 items, got %d", len(hashTags))
	}

	// 测试Tweet方法
	tweet := faker.Tweet()
	if tweet == "" {
		t.Error("Tweet() should not return empty string")
	}

	// 测试Review方法
	review := faker.Review()
	if review == "" {
		t.Error("Review() should not return empty string")
	}
}

// TestBatchTextFunctionsAdditional 测试批量文本生成函数
func TestBatchTextFunctionsAdditional(t *testing.T) {
	faker := New()

	// 测试BatchWords方法
	words := faker.BatchWords(5)
	if len(words) != 5 {
		t.Errorf("BatchWords(5) should return 5 items, got %d", len(words))
	}

	// 测试BatchSentences方法
	sentences := faker.BatchSentences(5)
	if len(sentences) != 5 {
		t.Errorf("BatchSentences(5) should return 5 items, got %d", len(sentences))
	}

	// 测试BatchParagraphs方法
	paragraphs := faker.BatchParagraphs(5)
	if len(paragraphs) != 5 {
		t.Errorf("BatchParagraphs(5) should return 5 items, got %d", len(paragraphs))
	}

	// 测试BatchTitles方法
	titles := faker.BatchTitles(5)
	if len(titles) != 5 {
		t.Errorf("BatchTitles(5) should return 5 items, got %d", len(titles))
	}
}
