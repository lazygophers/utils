//go:build (lang_es || lang_all) && (country_all || country_europe || country_lu || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLuxembourg.RegisterName(xlanguage.Spanish, "Luxemburgo")
	dataLuxembourg.RegisterOfficialName(xlanguage.Spanish, "Gran Ducado de Luxemburgo")
	dataLuxembourg.RegisterCapital(xlanguage.Spanish, "Luxemburgo")
}
