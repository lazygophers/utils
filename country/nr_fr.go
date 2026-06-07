//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNauru.RegisterName(xlanguage.French, "Nauru")
	dataNauru.RegisterOfficialName(xlanguage.French, "République de Nauru")
	dataNauru.RegisterCapital(xlanguage.French, "Yaren")
}
