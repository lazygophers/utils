//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_so)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSomalia.RegisterName(xlanguage.MustParse("zh-Hant"), "索馬利亞")
	dataSomalia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "索馬利亞聯邦共和國")
	dataSomalia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "摩加迪休")
}
