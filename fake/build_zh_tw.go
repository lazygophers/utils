//go:build fake_zh_tw

package fake

// 仅繁体中文支持的构建标签
// 使用: go build -tags fake_zh_tw

var supportedLanguagesOverride = []Language{LanguageChineseTraditional}
var defaultLanguageOverride = LanguageChineseTraditional

func init() {
	// 重写默认语言设置
	defaultFaker = New(WithLanguage(LanguageChineseTraditional))
	
	// 可选：预加载繁体中文数据
	getDataManager().PreloadData(LanguageChineseTraditional)
}

// GetSupportedLanguagesOverride 重写支持的语言列表（仅繁体中文）
func GetSupportedLanguagesOverride() []Language {
	return supportedLanguagesOverride
}