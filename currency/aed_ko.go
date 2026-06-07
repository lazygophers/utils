//go:build (lang_ko || lang_all) && (country_ae || country_all || country_asia || country_western_asia || currency_aed || currency_all)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	AED.RegisterName(xlanguage.Korean, "아랍에미리트 디르함")
}
