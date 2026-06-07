//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedKingdom.RegisterName(xlanguage.MustParse("zh-Hant"), "英國")
	dataUnitedKingdom.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "大不列顛及北愛爾蘭聯合王國")
	dataUnitedKingdom.RegisterCapital(xlanguage.MustParse("zh-Hant"), "倫敦")
}
