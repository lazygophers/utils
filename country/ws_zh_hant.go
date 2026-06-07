//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSamoa.RegisterName(xlanguage.MustParse("zh-Hant"), "薩摩亞")
	dataSamoa.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "薩摩亞獨立國")
	dataSamoa.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿庇亞")
}
