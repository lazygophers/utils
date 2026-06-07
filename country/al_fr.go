//go:build (lang_fr || lang_all) && (country_al || country_all || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlbania.RegisterName(xlanguage.French, "Albanie")
	dataAlbania.RegisterOfficialName(xlanguage.French, "République d'Albanie")
	dataAlbania.RegisterCapital(xlanguage.French, "Tirana")
}
