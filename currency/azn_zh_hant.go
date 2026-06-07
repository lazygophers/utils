//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_az || country_western_asia || currency_all || currency_azn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	AZN.RegisterName(xlanguage.MustParse("zh-Hant"), "亞塞拜然馬納特")
}
