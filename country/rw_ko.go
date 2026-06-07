//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_rw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRwanda.RegisterName(xlanguage.Korean, "르완다")
	dataRwanda.RegisterOfficialName(xlanguage.Korean, "르완다 공화국")
	dataRwanda.RegisterCapital(xlanguage.Korean, "키갈리")
}
