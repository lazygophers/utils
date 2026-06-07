//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAmericanSamoa.RegisterName(xlanguage.MustParse("zh-Hant"), "美屬薩摩亞")
	dataAmericanSamoa.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "美屬薩摩亞")
	dataAmericanSamoa.RegisterCapital(xlanguage.MustParse("zh-Hant"), "巴哥巴哥")
}
