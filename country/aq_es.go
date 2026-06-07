//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntarctica.RegisterName(xlanguage.Spanish, "Antártida")
	dataAntarctica.RegisterOfficialName(xlanguage.Spanish, "Antártida")
}
