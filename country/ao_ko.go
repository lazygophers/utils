//go:build (lang_ko || lang_all) && (country_africa || country_all || country_ao || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAngola.RegisterName(xlanguage.Korean, "앙골라")
	dataAngola.RegisterOfficialName(xlanguage.Korean, "앙골라 공화국")
	dataAngola.RegisterCapital(xlanguage.Korean, "루안다")
}
