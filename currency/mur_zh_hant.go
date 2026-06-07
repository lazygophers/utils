//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_mu || currency_all || currency_mur)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mur.RegisterName(xlanguage.MustParse("zh-Hant"), "模里西斯盧比")
}
