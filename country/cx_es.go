//go:build (lang_es || lang_all) && (country_all || country_australia_and_new_zealand || country_cx || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChristmasIsland.RegisterName(xlanguage.Spanish, "Isla de Navidad")
	dataChristmasIsland.RegisterOfficialName(xlanguage.Spanish, "Territorio de la Isla de Navidad")
	dataChristmasIsland.RegisterCapital(xlanguage.Spanish, "Flying Fish Cove")
}
