//go:build (lang_ko || lang_all) && (country_all || country_asia || country_om || country_western_asia || currency_all || currency_omr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Omr.RegisterName(xlanguage.Korean, "오만 리알")
}
