//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalta.RegisterName(xlanguage.French, "Malte")
	dataMalta.RegisterOfficialName(xlanguage.French, "République de Malte")
	dataMalta.RegisterCapital(xlanguage.French, "La Valette")
}
