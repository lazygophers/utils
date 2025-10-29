package fake

import (
	"fmt"
	"strings"

	"github.com/lazygophers/utils/randx"
)

// Word 生成单词
func (f *Faker) Word() string {

	values, weights, err := getWeightedItems(f.language, "texts", "lorem")
	if err != nil {
		// 回退到默认单词列表
		return f.getDefaultWord()
	}

	return randx.WeightedChoose(values, weights)
}

func (f *Faker) getDefaultWord() string {
	words := []string{
		"lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit",
		"sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore",
		"magna", "aliqua", "enim", "ad", "minim", "veniam", "quis", "nostrud",
		"exercitation", "ullamco", "laboris", "nisi", "aliquip", "ex", "ea", "commodo",
		"consequat", "duis", "aute", "irure", "in", "reprehenderit", "voluptate",
		"velit", "esse", "cillum", "fugiat", "nulla", "pariatur", "excepteur", "sint",
		"occaecat", "cupidatat", "non", "proident", "sunt", "culpa", "qui", "officia",
		"deserunt", "mollit", "anim", "id", "est", "laborum",
	}
	return randx.Choose(words)
}

// Words 生成多个单词
func (f *Faker) Words(count int) []string {

	if count <= 0 {
		return []string{}
	}

	words := make([]string, count)
	for i := 0; i < count; i++ {
		words[i] = f.Word()
	}

	return words
}

// Sentence 生成句子
func (f *Faker) Sentence() string {

	wordCount := randx.Intn(10) + 6 // 6-15个单词
	words := f.Words(wordCount)

	// 首字母大写
	if len(words) > 0 {
		words[0] = strings.Title(words[0])
	}

	sentence := strings.Join(words, " ")

	// 添加标点符号
	punctuation := []string{".", ".", ".", ".", "!", "?"}
	sentence += randx.Choose(punctuation)

	return sentence
}

// Sentences 生成多个句子
func (f *Faker) Sentences(count int) []string {
	f.incrementCallOut()

	if count <= 0 {
		return []string{}
	}

	sentences := make([]string, count)
	for i := 0; i < count; i++ {
		sentences[i] = f.Sentence()
	}

	return sentences
}

// Paragraph 生成段落
func (f *Faker) Paragraph() string {

	sentenceCount := randx.Intn(5) + 3 // 3-7个句子
	sentences := f.Sentences(sentenceCount)

	return strings.Join(sentences, " ")
}

// Paragraphs 生成多个段落
func (f *Faker) Paragraphs(count int) []string {

	if count <= 0 {
		return []string{}
	}

	paragraphs := make([]string, count)
	for i := 0; i < count; i++ {
		paragraphs[i] = f.Paragraph()
	}

	return paragraphs
}

// Text 生成指定长度的文本（字符数）
func (f *Faker) Text(maxChars int) string {

	if maxChars <= 0 {
		return ""
	}

	var text strings.Builder

	for text.Len() < maxChars {
		sentence := f.Sentence()

		// 检查添加这个句子是否会超出长度限制
		if text.Len()+len(sentence)+1 > maxChars {
			// 如果会超出，只添加部分内容
			remaining := maxChars - text.Len()
			if remaining > 0 {
				if text.Len() > 0 {
					text.WriteString(" ")
					remaining--
				}
				if remaining > 0 {
					text.WriteString(sentence[:remaining])
				}
			}
			break
		}

		if text.Len() > 0 {
			text.WriteString(" ")
		}
		text.WriteString(sentence)
	}

	return text.String()
}

// Title 生成标题
func (f *Faker) Title() string {

	wordCount := randx.Intn(5) + 2 // 2-6个单词
	words := f.Words(wordCount)

	// 标题化每个单词（首字母大写）
	for i, word := range words {
		words[i] = strings.Title(word)
	}

	return strings.Join(words, " ")
}

// Quote 生成引用
func (f *Faker) Quote() string {

	quote := f.Sentence()

	// 移除结尾的标点符号
	quote = strings.TrimSuffix(quote, ".")
	quote = strings.TrimSuffix(quote, "!")
	quote = strings.TrimSuffix(quote, "?")

	return fmt.Sprintf("\"%s\"", quote)
}

// Lorem 生成Lorem Ipsum文本
func (f *Faker) Lorem() string {

	return "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
}

// LoremWords 生成Lorem Ipsum单词
func (f *Faker) LoremWords(count int) string {

	words := f.Words(count)
	return strings.Join(words, " ")
}

// LoremSentences 生成Lorem Ipsum句子
func (f *Faker) LoremSentences(count int) string {

	sentences := f.Sentences(count)
	return strings.Join(sentences, " ")
}

// LoremParagraphs 生成Lorem Ipsum段落
func (f *Faker) LoremParagraphs(count int) string {

	paragraphs := f.Paragraphs(count)
	return strings.Join(paragraphs, "\n\n")
}

// Article 生成文章
func (f *Faker) Article() string {

	title := f.Title()
	paragraphCount := randx.Intn(5) + 3 // 3-7个段落
	paragraphs := f.Paragraphs(paragraphCount)

	var article strings.Builder
	article.WriteString(title)
	article.WriteString("\n\n")

	for i, paragraph := range paragraphs {
		if i > 0 {
			article.WriteString("\n\n")
		}
		article.WriteString(paragraph)
	}

	return article.String()
}

// Slug 生成URL slug
func (f *Faker) Slug() string {

	wordCount := randx.Intn(4) + 2 // 2-5个单词
	words := f.Words(wordCount)

	// 转换为小写并用连字符连接
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return strings.Join(words, "-")
}

// HashTag 生成标签
func (f *Faker) HashTag() string {

	word := f.Word()
	return fmt.Sprintf("#%s", strings.ToLower(word))
}

// HashTags 生成多个标签
func (f *Faker) HashTags(count int) []string {

	if count <= 0 {
		return []string{}
	}

	tags := make([]string, count)
	for i := 0; i < count; i++ {
		tags[i] = f.HashTag()
	}

	return tags
}

// Tweet 生成类似推特的短消息
func (f *Faker) Tweet() string {

	// 推特限制280字符
	text := f.Text(200) // 留出空间给标签和链接

	// 30% 概率添加标签
	if randx.Float32() < 0.3 {
		tags := f.HashTags(randx.Intn(3) + 1)
		text += " " + strings.Join(tags, " ")
	}

	// 20% 概率添加链接
	if randx.Float32() < 0.2 {
		text += " " + f.URL()
	}

	// 确保不超过280字符
	if len(text) > 280 {
		text = text[:277] + "..."
	}

	return text
}

// Review 生成评论/评价
func (f *Faker) Review() string {

	// 评价通常比较简短
	sentenceCount := randx.Intn(4) + 1 // 1-4个句子

	var review strings.Builder

	// 可能以评价开头
	ratings := []string{
		"Great product!",
		"Excellent service!",
		"Good quality.",
		"Not bad.",
		"Could be better.",
		"Disappointing.",
		"Amazing!",
		"Wonderful experience.",
		"Highly recommended!",
		"Will buy again.",
	}

	if randx.Float32() < 0.3 {
		review.WriteString(randx.Choose(ratings))
		review.WriteString(" ")
		sentenceCount--
	}

	sentences := f.Sentences(sentenceCount)
	review.WriteString(strings.Join(sentences, " "))

	return review.String()
}

// 批量生成函数
func (f *Faker) BatchWords(count int) []string {
	return f.batchGenerate(count, f.Word)
}

func (f *Faker) BatchSentences(count int) []string {
	return f.batchGenerate(count, f.Sentence)
}

func (f *Faker) BatchParagraphs(count int) []string {
	return f.batchGenerate(count, f.Paragraph)
}

func (f *Faker) BatchTitles(count int) []string {
	return f.batchGenerate(count, f.Title)
}

// 修复函数名拼写错误
func (f *Faker) incrementCallOut() {
}

// 全局便捷函数
func Word() string {
	return getDefaultFaker().Word()
}

func Words(count int) []string {
	return getDefaultFaker().Words(count)
}

func Sentence() string {
	return getDefaultFaker().Sentence()
}

func Sentences(count int) []string {
	return getDefaultFaker().Sentences(count)
}

func Paragraph() string {
	return getDefaultFaker().Paragraph()
}

func Paragraphs(count int) []string {
	return getDefaultFaker().Paragraphs(count)
}

func Text(maxChars int) string {
	return getDefaultFaker().Text(maxChars)
}

func Title() string {
	return getDefaultFaker().Title()
}

func Quote() string {
	return getDefaultFaker().Quote()
}

func Lorem() string {
	return getDefaultFaker().Lorem()
}

func LoremWords(count int) string {
	return getDefaultFaker().LoremWords(count)
}

func LoremSentences(count int) string {
	return getDefaultFaker().LoremSentences(count)
}

func LoremParagraphs(count int) string {
	return getDefaultFaker().LoremParagraphs(count)
}

func Article() string {
	return getDefaultFaker().Article()
}

func Slug() string {
	return getDefaultFaker().Slug()
}

func HashTag() string {
	return getDefaultFaker().HashTag()
}

func HashTags(count int) []string {
	return getDefaultFaker().HashTags(count)
}

func Tweet() string {
	return getDefaultFaker().Tweet()
}

func Review() string {
	return getDefaultFaker().Review()
}
