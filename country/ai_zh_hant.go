//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAnguilla.RegisterName(xlanguage.MustParse("zh-Hant"), "安圭拉")
	dataAnguilla.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "安圭拉")
	dataAnguilla.RegisterCapital(xlanguage.MustParse("zh-Hant"), "谷地")
}
