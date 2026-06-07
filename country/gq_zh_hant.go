//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_gq || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEquatorialGuinea.RegisterName(xlanguage.MustParse("zh-Hant"), "赤道幾內亞")
	dataEquatorialGuinea.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "赤道幾內亞共和國")
	dataEquatorialGuinea.RegisterCapital(xlanguage.MustParse("zh-Hant"), "馬拉博")
}
