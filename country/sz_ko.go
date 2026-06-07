//go:build (lang_ko || lang_all) && (country_africa || country_all || country_southern_africa || country_sz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEswatini.RegisterName(xlanguage.Korean, "에스와티니")
	dataEswatini.RegisterOfficialName(xlanguage.Korean, "에스와티니 왕국")
	dataEswatini.RegisterCapital(xlanguage.Korean, "음바바네")
}
