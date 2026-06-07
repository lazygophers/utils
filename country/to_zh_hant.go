//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTonga.RegisterName(xlanguage.MustParse("zh-Hant"), "東加")
	dataTonga.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "東加王國")
	dataTonga.RegisterCapital(xlanguage.MustParse("zh-Hant"), "努瓜婁發")
}
