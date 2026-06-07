//go:build (lang_ko || lang_all) && (country_all || country_americas || country_south_america || country_ve)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVenezuela.RegisterName(xlanguage.Korean, "베네수엘라")
	dataVenezuela.RegisterOfficialName(xlanguage.Korean, "베네수엘라 볼리바르 공화국")
	dataVenezuela.RegisterCapital(xlanguage.Korean, "카라카스")
}
