//go:build (lang_fr || lang_all) && (country_all || country_ax || country_europe || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlandIslands.RegisterName(xlanguage.French, "Åland")
	dataAlandIslands.RegisterOfficialName(xlanguage.French, "Åland")
	dataAlandIslands.RegisterCapital(xlanguage.French, "Mariehamn")
}
