//go:build (lang_fr || lang_all) && (country_all || country_europe || country_si || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovenia.RegisterName(xlanguage.French, "Slovénie")
	dataSlovenia.RegisterOfficialName(xlanguage.French, "République de Slovénie")
	dataSlovenia.RegisterCapital(xlanguage.French, "Ljubljana")
}
