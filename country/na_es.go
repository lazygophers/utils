//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNamibia.RegisterName(xlanguage.Spanish, "Namibia")
	dataNamibia.RegisterOfficialName(xlanguage.Spanish, "República de Namibia")
	dataNamibia.RegisterCapital(xlanguage.Spanish, "Windhoek")
}
