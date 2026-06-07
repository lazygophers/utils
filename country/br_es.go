//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrazil.RegisterName(xlanguage.Spanish, "Brasil")
	dataBrazil.RegisterOfficialName(xlanguage.Spanish, "República Federativa del Brasil")
	dataBrazil.RegisterCapital(xlanguage.Spanish, "Brasilia")
}
