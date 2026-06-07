//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHaiti.RegisterName(xlanguage.Spanish, "Haití")
	dataHaiti.RegisterOfficialName(xlanguage.Spanish, "República de Haití")
	dataHaiti.RegisterCapital(xlanguage.Spanish, "Puerto Príncipe")
}
