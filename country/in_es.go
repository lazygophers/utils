//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndia.RegisterName(xlanguage.Spanish, "India")
	dataIndia.RegisterOfficialName(xlanguage.Spanish, "República de la India")
	dataIndia.RegisterCapital(xlanguage.Spanish, "Nueva Delhi")
}
