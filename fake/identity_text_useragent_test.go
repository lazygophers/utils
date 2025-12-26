package fake

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentityMethods(t *testing.T) {
	faker := New()

	t.Run("ChineseIDNumber", func(t *testing.T) {
		id := faker.ChineseIDNumber()
		assert.NotEmpty(t, id)
		assert.Len(t, id, 18)
	})

	t.Run("DriversLicense", func(t *testing.T) {
		license := faker.DriversLicense()
		assert.NotEmpty(t, license)
		assert.GreaterOrEqual(t, len(license), 10)
	})

	t.Run("Passport", func(t *testing.T) {
		passport := faker.Passport()
		assert.NotEmpty(t, passport)
		assert.GreaterOrEqual(t, len(passport), 8)
	})

	t.Run("IdentityDoc", func(t *testing.T) {
		doc := faker.IdentityDoc()
		assert.NotEmpty(t, doc)
	})

	t.Run("CreditCardNumber", func(t *testing.T) {
		card := faker.CreditCardNumber()
		assert.NotEmpty(t, card)
		assert.GreaterOrEqual(t, len(card), 13)
		assert.LessOrEqual(t, len(card), 19)
	})

	t.Run("SSN", func(t *testing.T) {
		ssn := faker.SSN()
		assert.NotEmpty(t, ssn)
		assert.Len(t, ssn, 11)
		assert.Contains(t, ssn, "-")
	})
}

func TestTextMethods(t *testing.T) {
	faker := New()

	t.Run("Words", func(t *testing.T) {
		words := faker.Words(5)
		assert.Len(t, words, 5)
		for _, word := range words {
			assert.NotEmpty(t, word)
		}
	})

	t.Run("Sentences", func(t *testing.T) {
		sentences := faker.Sentences(3)
		assert.Len(t, sentences, 3)
		for _, sentence := range sentences {
			assert.NotEmpty(t, sentence)
		}
	})

	t.Run("Paragraphs", func(t *testing.T) {
		paragraphs := faker.Paragraphs(2)
		assert.Len(t, paragraphs, 2)
		for _, paragraph := range paragraphs {
			assert.NotEmpty(t, paragraph)
		}
	})

	t.Run("Text", func(t *testing.T) {
		text := faker.Text(100)
		assert.NotEmpty(t, text)
		assert.GreaterOrEqual(t, len(text), 50)
	})

	t.Run("Quote", func(t *testing.T) {
		quote := faker.Quote()
		assert.NotEmpty(t, quote)
	})

	t.Run("Lorem", func(t *testing.T) {
		lorem := faker.Lorem()
		assert.NotEmpty(t, lorem)
	})

	t.Run("LoremWords", func(t *testing.T) {
		words := faker.LoremWords(5)
		assert.NotEmpty(t, words)
	})

	t.Run("LoremSentences", func(t *testing.T) {
		sentences := faker.LoremSentences(3)
		assert.NotEmpty(t, sentences)
	})

	t.Run("LoremParagraphs", func(t *testing.T) {
		paragraphs := faker.LoremParagraphs(2)
		assert.NotEmpty(t, paragraphs)
	})

	t.Run("Article", func(t *testing.T) {
		article := faker.Article()
		assert.NotEmpty(t, article)
	})

	t.Run("Slug", func(t *testing.T) {
		slug := faker.Slug()
		assert.NotEmpty(t, slug)
		assert.NotContains(t, slug, " ")
	})

	t.Run("HashTag", func(t *testing.T) {
		tag := faker.HashTag()
		assert.NotEmpty(t, tag)
		assert.Contains(t, tag, "#")
	})

	t.Run("HashTags", func(t *testing.T) {
		tags := faker.HashTags(5)
		assert.Len(t, tags, 5)
		for _, tag := range tags {
			assert.Contains(t, tag, "#")
		}
	})

	t.Run("Tweet", func(t *testing.T) {
		tweet := faker.Tweet()
		assert.NotEmpty(t, tweet)
		assert.LessOrEqual(t, len(tweet), 280)
	})

	t.Run("Review", func(t *testing.T) {
		review := faker.Review()
		assert.NotEmpty(t, review)
	})
}

func TestUserAgentMethods(t *testing.T) {
	faker := New()

	t.Run("UserAgentFor", func(t *testing.T) {
		ua := faker.UserAgentFor("Chrome")
		assert.NotEmpty(t, ua)
		assert.True(t, strings.Contains(ua, "Chrome") || strings.Contains(ua, "chrome"))
	})

	t.Run("UserAgentForPlatform", func(t *testing.T) {
		ua := faker.UserAgentForPlatform("windows")
		assert.NotEmpty(t, ua)
	})

	t.Run("UserAgentForDevice", func(t *testing.T) {
		ua := faker.UserAgentForDevice("desktop")
		assert.NotEmpty(t, ua)
	})

	t.Run("ChromeUserAgent", func(t *testing.T) {
		ua := faker.ChromeUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Chrome")
	})

	t.Run("FirefoxUserAgent", func(t *testing.T) {
		ua := faker.FirefoxUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Firefox")
	})

	t.Run("SafariUserAgent", func(t *testing.T) {
		ua := faker.SafariUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Safari")
	})

	t.Run("EdgeUserAgent", func(t *testing.T) {
		ua := faker.EdgeUserAgent()
		assert.NotEmpty(t, ua)
	})

	t.Run("AndroidUserAgent", func(t *testing.T) {
		ua := faker.AndroidUserAgent()
		assert.NotEmpty(t, ua)
		assert.Contains(t, ua, "Android")
	})

	t.Run("IOSUserAgent", func(t *testing.T) {
		ua := faker.IOSUserAgent()
		assert.NotEmpty(t, ua)
		hasIPhone := strings.Contains(ua, "iPhone")
		hasIPad := strings.Contains(ua, "iPad")
		assert.True(t, hasIPhone || hasIPad)
	})
}
