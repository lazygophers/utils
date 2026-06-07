//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCzechia.RegisterName(xlanguage.Spanish, "Chequia")
	dataCzechia.RegisterOfficialName(xlanguage.Spanish, "República Checa")
	dataCzechia.RegisterCapital(xlanguage.Spanish, "Praga")
}
