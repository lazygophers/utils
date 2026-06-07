//go:build (lang_zh_hant || lang_all) && (country_all || country_au || country_australia_and_new_zealand || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustralia.RegisterName(xlanguage.MustParse("zh-Hant"), "澳大利亞")
	dataAustralia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "澳大利亞聯邦")
	dataAustralia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "坎培拉")
}
