//go:build (lang_ko || lang_all) && (country_africa || country_all || country_sl || country_western_africa || currency_all || currency_sll)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sll.RegisterName(xlanguage.Korean, "시에라리온 레온")
}
