//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTuvalu.RegisterName(xlanguage.MustParse("zh-Hant"), "吐瓦魯")
	dataTuvalu.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "吐瓦魯")
	dataTuvalu.RegisterCapital(xlanguage.MustParse("zh-Hant"), "富那富提")
}
