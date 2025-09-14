//go:build fake_multi

package fake

// 多语言支持的构建标签（仅包含最常用的语言）
// 使用: go build -tags fake_multi

var supportedLanguagesOverride = []Language{
	LanguageEnglish,
	LanguageChineseSimplified,
	LanguageChineseTraditional,
}

func init() {
	// 预加载常用语言数据
	for _, lang := range supportedLanguagesOverride {
		getDataManager().PreloadData(lang)
	}
}

// GetSupportedLanguagesOverride 重写支持的语言列表（多语言）
func GetSupportedLanguagesOverride() []Language {
	return supportedLanguagesOverride
}