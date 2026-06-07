//go:build (lang_es || lang_all) && (country_africa || country_all || country_gm || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGambia.RegisterName(xlanguage.Spanish, "Gambia")
	dataGambia.RegisterOfficialName(xlanguage.Spanish, "República de Gambia")
	dataGambia.RegisterCapital(xlanguage.Spanish, "Banjul")
}
