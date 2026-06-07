//go:build (lang_fr || lang_all) && (country_ad || country_all || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAndorra.RegisterName(xlanguage.French, "Andorre")
	dataAndorra.RegisterOfficialName(xlanguage.French, "Principauté d'Andorre")
	dataAndorra.RegisterCapital(xlanguage.French, "Andorre-la-Vieille")
}
