//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedStates.RegisterName(xlanguage.MustParse("zh-Hant"), "美國")
	dataUnitedStates.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "美利堅合眾國")
	dataUnitedStates.RegisterCapital(xlanguage.MustParse("zh-Hant"), "華盛頓哥倫比亞特區")
}
