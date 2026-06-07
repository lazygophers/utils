//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCocosKeelingIslands.RegisterName(xlanguage.MustParse("zh-Hant"), "科科斯（基林）群島")
	dataCocosKeelingIslands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "科科斯（基林）群島")
	dataCocosKeelingIslands.RegisterCapital(xlanguage.MustParse("zh-Hant"), "西島")
}
