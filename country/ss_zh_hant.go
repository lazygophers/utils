//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthSudan.RegisterName(xlanguage.MustParse("zh-Hant"), "南蘇丹")
	dataSouthSudan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "南蘇丹共和國")
	dataSouthSudan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "朱巴")
}
