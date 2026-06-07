//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSweden.RegisterName(xlanguage.Spanish, "Suecia")
	dataSweden.RegisterOfficialName(xlanguage.Spanish, "Reino de Suecia")
	dataSweden.RegisterCapital(xlanguage.Spanish, "Estocolmo")
}
