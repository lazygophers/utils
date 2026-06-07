//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFinland.RegisterName(xlanguage.Spanish, "Finlandia")
	dataFinland.RegisterOfficialName(xlanguage.Spanish, "República de Finlandia")
	dataFinland.RegisterCapital(xlanguage.Spanish, "Helsinki")
}
