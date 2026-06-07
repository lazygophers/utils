//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiger.RegisterName(xlanguage.Spanish, "Níger")
	dataNiger.RegisterOfficialName(xlanguage.Spanish, "República de Níger")
	dataNiger.RegisterCapital(xlanguage.Spanish, "Niamey")
}
