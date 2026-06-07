package fake

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// loremWords is a classic Latin lorem ipsum pool used by English text
// generators. It intentionally avoids punctuation; callers are responsible
// for capitalisation and sentence terminators.
var loremWords = []string{
	"lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing",
	"elit", "sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore",
	"et", "dolore", "magna", "aliqua", "enim", "ad", "minim", "veniam",
	"quis", "nostrud", "exercitation", "ullamco", "laboris", "nisi",
	"aliquip", "ex", "ea", "commodo", "consequat", "duis", "aute", "irure",
	"in", "reprehenderit", "voluptate", "velit", "esse", "cillum", "eu",
	"fugiat", "nulla", "pariatur", "excepteur", "sint", "occaecat",
	"cupidatat", "non", "proident", "sunt", "culpa", "qui", "officia",
	"deserunt", "mollit", "anim", "id", "est", "laborum", "at", "vero",
	"eos", "accusamus", "iusto", "odio", "dignissimos", "ducimus",
	"blanditiis", "praesentium", "voluptatum", "deleniti", "atque",
	"corrupti", "quos", "dolores", "quas", "molestias", "excepturi",
	"obcaecati", "cupiditate", "provident", "similique", "mollitia",
	"animi", "dolorum", "fuga", "harum", "quidem", "rerum", "facilis",
	"expedita", "distinctio", "nam", "libero", "tempore", "cum", "soluta",
	"nobis", "eligendi", "optio", "cumque", "impedit", "minus", "quod",
	"maxime", "placeat", "facere", "possimus", "omnis", "voluptas",
	"assumenda", "repellendus", "temporibus", "autem", "quibusdam",
	"officiis", "debitis", "necessitatibus", "saepe", "eveniet",
	"aspernatur",
}

// chineseWords is a pool of common Mandarin bi-character compounds used by
// Chinese text generators. Each entry is a self-contained word without
// punctuation so callers can freely concatenate them.
var chineseWords = []string{
	"今天", "明天", "昨天", "早晨", "中午", "傍晚", "夜晚", "春风", "夏雨", "秋叶",
	"冬雪", "山川", "河流", "草原", "森林", "海洋", "湖泊", "城市", "乡村", "田野",
	"道路", "桥梁", "高楼", "房屋", "花园", "庭院", "公园", "学校", "医院", "商店",
	"餐厅", "咖啡", "茶水", "美食", "旅行", "工作", "学习", "阅读", "写作", "思考",
	"讨论", "决定", "计划", "目标", "梦想", "希望", "友谊", "家庭", "亲人", "朋友",
	"邻居", "同事", "老师", "学生", "医生", "工程师", "艺术家", "音乐家", "画家", "作家",
	"演员", "导演", "设计师", "程序员", "经理", "老板", "客户", "用户", "产品", "服务",
	"质量", "效率", "创新", "技术", "科学", "文化", "历史", "哲学", "经济", "政治",
	"社会", "自然", "生命", "时间", "空间",
}

// Word returns a single lowercase lorem ipsum word.
func (f *Faker) Word() string {
	return f.pickString(loremWords)
}

// Words returns a slice of n lowercase lorem ipsum words. When n is
// non-positive an empty slice is returned.
func (f *Faker) Words(n int) []string {
	if n <= 0 {
		return []string{}
	}
	out := make([]string, n)
	for i := 0; i < n; i++ {
		out[i] = f.pickString(loremWords)
	}
	return out
}

// Sentence returns a single lorem ipsum sentence containing wordCount
// words. When wordCount is non-positive a length in [6, 12] is chosen at
// random. The first word is capitalised and the sentence terminates with a
// period.
func (f *Faker) Sentence(wordCount int) string {
	if wordCount <= 0 {
		wordCount = 6 + f.intN(7)
	}
	var b strings.Builder
	b.Grow(wordCount * 8)
	for i := 0; i < wordCount; i++ {
		w := f.pickString(loremWords)
		if i == 0 {
			r, size := utf8.DecodeRuneInString(w)
			b.WriteRune(unicode.ToUpper(r))
			b.WriteString(w[size:])
		} else {
			b.WriteByte(' ')
			b.WriteString(w)
		}
	}
	b.WriteByte('.')
	return b.String()
}

// Paragraph returns sentenceCount lorem ipsum sentences joined by a single
// space. When sentenceCount is non-positive a count in [3, 6] is chosen at
// random.
func (f *Faker) Paragraph(sentenceCount int) string {
	if sentenceCount <= 0 {
		sentenceCount = 3 + f.intN(4)
	}
	parts := make([]string, sentenceCount)
	for i := 0; i < sentenceCount; i++ {
		parts[i] = f.Sentence(0)
	}
	return strings.Join(parts, " ")
}

// ChineseWord returns a single common Mandarin compound word.
func (f *Faker) ChineseWord() string {
	return f.pickString(chineseWords)
}

// ChineseWords returns n Mandarin words. When n is non-positive an empty
// slice is returned.
func (f *Faker) ChineseWords(n int) []string {
	if n <= 0 {
		return []string{}
	}
	out := make([]string, n)
	for i := 0; i < n; i++ {
		out[i] = f.pickString(chineseWords)
	}
	return out
}

// ChineseSentence returns a Mandarin sentence with approximately charCount
// Han characters, terminated by the full-width period "。". When charCount
// is non-positive a target length in [15, 30] characters is chosen at
// random. The actual length may overshoot by up to one word because words
// are appended atomically.
func (f *Faker) ChineseSentence(charCount int) string {
	if charCount <= 0 {
		charCount = 15 + f.intN(16)
	}
	var b strings.Builder
	b.Grow(charCount * 3)
	count := 0
	for count < charCount {
		w := f.pickString(chineseWords)
		b.WriteString(w)
		count += utf8.RuneCountInString(w)
	}
	b.WriteRune('。')
	return b.String()
}

// ChineseParagraph returns sentenceCount Mandarin sentences concatenated
// without separators (sentences already terminate with "。"). When
// sentenceCount is non-positive a count in [3, 6] is chosen at random.
func (f *Faker) ChineseParagraph(sentenceCount int) string {
	if sentenceCount <= 0 {
		sentenceCount = 3 + f.intN(4)
	}
	var b strings.Builder
	for i := 0; i < sentenceCount; i++ {
		b.WriteString(f.ChineseSentence(0))
	}
	return b.String()
}
