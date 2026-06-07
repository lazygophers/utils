//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedArabEmirates.RegisterName(xlanguage.MustParse("zh-Hant"), "阿拉伯聯合大公國")
	dataUnitedArabEmirates.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "阿拉伯聯合大公國")
	dataUnitedArabEmirates.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿布達比")
}
