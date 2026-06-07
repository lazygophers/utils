//go:build (lang_fr || lang_all) && (country_all || country_eastern_europe || country_europe || country_pl)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPoland.RegisterName(xlanguage.French, "Pologne")
	dataPoland.RegisterOfficialName(xlanguage.French, "République de Pologne")
	dataPoland.RegisterCapital(xlanguage.French, "Varsovie")
}
