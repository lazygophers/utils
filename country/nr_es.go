//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNauru.RegisterName(xlanguage.Spanish, "Nauru")
	dataNauru.RegisterOfficialName(xlanguage.Spanish, "República de Nauru")
	dataNauru.RegisterCapital(xlanguage.Spanish, "Yaren")
}
