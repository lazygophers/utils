//go:build (lang_es || lang_all) && (country_all || country_europe || country_si || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovenia.RegisterName(xlanguage.Spanish, "Eslovenia")
	dataSlovenia.RegisterOfficialName(xlanguage.Spanish, "República de Eslovenia")
	dataSlovenia.RegisterCapital(xlanguage.Spanish, "Liubliana")
}
