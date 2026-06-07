//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSintMaarten.RegisterName(xlanguage.MustParse("zh-Hant"), "荷屬聖馬丁")
	dataSintMaarten.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "荷屬聖馬丁")
	dataSintMaarten.RegisterCapital(xlanguage.MustParse("zh-Hant"), "菲利普斯堡")
}
