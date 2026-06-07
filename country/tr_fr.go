//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkey.RegisterName(xlanguage.French, "Turquie")
	dataTurkey.RegisterOfficialName(xlanguage.French, "République de Turquie")
	dataTurkey.RegisterCapital(xlanguage.French, "Ankara")
}
