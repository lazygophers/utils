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

func TestBatchCompanyNames(t *testing.T) {
	faker := New()

	companies := faker.BatchCompanyNames(5)
	assert.Len(t, companies, 5)

	for _, company := range companies {
		assert.NotEmpty(t, company)
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
