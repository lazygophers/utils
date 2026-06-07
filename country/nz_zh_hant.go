//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewZealand.RegisterName(xlanguage.MustParse("zh-Hant"), "紐西蘭")
	dataNewZealand.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "紐西蘭")
	dataNewZealand.RegisterCapital(xlanguage.MustParse("zh-Hant"), "威靈頓")
}
