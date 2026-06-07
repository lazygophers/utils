//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAndorra.RegisterName(xlanguage.Spanish, "Andorra")
	dataAndorra.RegisterOfficialName(xlanguage.Spanish, "Principado de Andorra")
	dataAndorra.RegisterCapital(xlanguage.Spanish, "Andorra la Vieja")
}
