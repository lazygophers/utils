//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurksAndCaicosIslands.RegisterName(xlanguage.Spanish, "Islas Turcas y Caicos")
	dataTurksAndCaicosIslands.RegisterOfficialName(xlanguage.Spanish, "Islas Turcas y Caicos")
	dataTurksAndCaicosIslands.RegisterCapital(xlanguage.Spanish, "Cockburn Town")
}
