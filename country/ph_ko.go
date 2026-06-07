//go:build (lang_ko || lang_all) && (country_all || country_asia || country_ph || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPhilippines.RegisterName(xlanguage.Korean, "필리핀")
	dataPhilippines.RegisterOfficialName(xlanguage.Korean, "필리핀 공화국")
	dataPhilippines.RegisterCapital(xlanguage.Korean, "마닐라")
}
