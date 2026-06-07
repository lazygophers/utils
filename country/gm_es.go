//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGambia.RegisterName(xlanguage.Spanish, "Gambia")
	dataGambia.RegisterOfficialName(xlanguage.Spanish, "República de Gambia")
	dataGambia.RegisterCapital(xlanguage.Spanish, "Banjul")
}
