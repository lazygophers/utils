//go:build (lang_zh_hant || lang_all) && (country_all || country_melanesia || country_oceania || country_vu)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVanuatu.RegisterName(xlanguage.MustParse("zh-Hant"), "萬那杜")
	dataVanuatu.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "萬那杜共和國")
	dataVanuatu.RegisterCapital(xlanguage.MustParse("zh-Hant"), "維拉港")
}
