//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNamibia.RegisterName(xlanguage.French, "Namibie")
	dataNamibia.RegisterOfficialName(xlanguage.French, "République de Namibie")
	dataNamibia.RegisterCapital(xlanguage.French, "Windhoek")
}
