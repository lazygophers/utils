//go:build (lang_ko || lang_all) && (country_all || country_asia || country_tr || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkey.RegisterName(xlanguage.Korean, "튀르키예")
	dataTurkey.RegisterOfficialName(xlanguage.Korean, "튀르키예 공화국")
	dataTurkey.RegisterCapital(xlanguage.Korean, "앙카라")
}
