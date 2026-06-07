//go:build (lang_fr || lang_all) && (country_all || country_europe || country_fi || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFinland.RegisterName(xlanguage.French, "Finlande")
	dataFinland.RegisterOfficialName(xlanguage.French, "République de Finlande")
	dataFinland.RegisterCapital(xlanguage.French, "Helsinki")
}
