//go:build (lang_ko || lang_all) && (country_all || country_asia || country_bt || country_southern_asia || currency_all || currency_btn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BTN.RegisterName(xlanguage.Korean, "부탄 눌탐")
}
