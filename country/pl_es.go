//go:build (lang_es || lang_all) && (country_all || country_eastern_europe || country_europe || country_pl)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPoland.RegisterName(xlanguage.Spanish, "Polonia")
	dataPoland.RegisterOfficialName(xlanguage.Spanish, "República de Polonia")
	dataPoland.RegisterCapital(xlanguage.Spanish, "Varsovia")
}
