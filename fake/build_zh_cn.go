//go:build fake_zh_cn

package fake

// 仅简体中文支持的构建标签
// 使用: go build -tags fake_zh_cn

var supportedLanguagesOverride = []Language{LanguageChineseSimplified}
var defaultLanguageOverride = LanguageChineseSimplified

func init() {
	// 重写默认语言设置
	defaultFaker = New(WithLanguage(LanguageChineseSimplified), WithCountry(CountryChina))
	
	// 可选：预加载中文数据
	getDataManager().PreloadData(LanguageChineseSimplified)
}

// GetSupportedLanguagesOverride 重写支持的语言列表（仅简体中文）
func GetSupportedLanguagesOverride() []Language {
	return supportedLanguagesOverride
}