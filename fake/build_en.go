//go:build fake_en

package fake

// 仅英语支持的构建标签
// 使用: go build -tags fake_en

var supportedLanguagesOverride = []Language{LanguageEnglish}
var defaultLanguageOverride = LanguageEnglish

func init() {
	// 重写默认语言设置
	defaultFaker = New(WithLanguage(LanguageEnglish))
	
	// 可选：预加载英语数据
	getDataManager().PreloadData(LanguageEnglish)
}

// GetSupportedLanguages 重写支持的语言列表（仅英语）
func GetSupportedLanguagesOverride() []Language {
	return supportedLanguagesOverride
}