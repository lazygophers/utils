//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGambia.RegisterName(xlanguage.French, "Gambie")
	dataGambia.RegisterOfficialName(xlanguage.French, "République de Gambie")
	dataGambia.RegisterCapital(xlanguage.French, "Banjul")
}
