//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLithuania.RegisterName(xlanguage.French, "Lituanie")
	dataLithuania.RegisterOfficialName(xlanguage.French, "République de Lituanie")
	dataLithuania.RegisterCapital(xlanguage.French, "Vilnius")
}
