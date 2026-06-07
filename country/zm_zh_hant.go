//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZambia.RegisterName(xlanguage.MustParse("zh-Hant"), "尚比亞")
	dataZambia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "尚比亞共和國")
	dataZambia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "路沙卡")
}
