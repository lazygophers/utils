//go:build (lang_zh_hant || lang_all) && (country_ae || country_all || country_asia || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedArabEmirates.RegisterName(xlanguage.MustParse("zh-Hant"), "阿拉伯聯合大公國")
	dataUnitedArabEmirates.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "阿拉伯聯合大公國")
	dataUnitedArabEmirates.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿布達比")
}
