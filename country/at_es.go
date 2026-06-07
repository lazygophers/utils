//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustria.RegisterName(xlanguage.Spanish, "Austria")
	dataAustria.RegisterOfficialName(xlanguage.Spanish, "República de Austria")
	dataAustria.RegisterCapital(xlanguage.Spanish, "Viena")
}
