//go:build (lang_fr || lang_all) && (country_all || country_asia || country_south_eastern_asia || country_th)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataThailand.RegisterName(xlanguage.French, "Thaïlande")
	dataThailand.RegisterOfficialName(xlanguage.French, "Royaume de Thaïlande")
	dataThailand.RegisterCapital(xlanguage.French, "Bangkok")
}
