//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGermany.RegisterName(xlanguage.Spanish, "Alemania")
	dataGermany.RegisterOfficialName(xlanguage.Spanish, "República Federal de Alemania")
	dataGermany.RegisterCapital(xlanguage.Spanish, "Berlín")
}
