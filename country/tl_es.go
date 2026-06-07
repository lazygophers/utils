//go:build (lang_es || lang_all) && (country_all || country_asia || country_south_eastern_asia || country_tl)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTimorLeste.RegisterName(xlanguage.Spanish, "Timor Oriental")
	dataTimorLeste.RegisterOfficialName(xlanguage.Spanish, "República Democrática de Timor Oriental")
	dataTimorLeste.RegisterCapital(xlanguage.Spanish, "Dili")
}
