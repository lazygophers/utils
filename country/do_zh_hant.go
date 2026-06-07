//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominicanRepublic.RegisterName(xlanguage.MustParse("zh-Hant"), "多明尼加")
	dataDominicanRepublic.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "多明尼加共和國")
	dataDominicanRepublic.RegisterCapital(xlanguage.MustParse("zh-Hant"), "聖多明哥")
}
