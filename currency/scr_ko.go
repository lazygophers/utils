//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_sc || currency_all || currency_scr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Scr.RegisterName(xlanguage.Korean, "세이셸 루피")
}
