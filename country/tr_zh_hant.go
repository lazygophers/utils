//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkey.RegisterName(xlanguage.MustParse("zh-Hant"), "土耳其")
	dataTurkey.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "土耳其共和國")
	dataTurkey.RegisterCapital(xlanguage.MustParse("zh-Hant"), "安卡拉")
}
