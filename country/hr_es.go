//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCroatia.RegisterName(xlanguage.Spanish, "Croacia")
	dataCroatia.RegisterOfficialName(xlanguage.Spanish, "República de Croacia")
	dataCroatia.RegisterCapital(xlanguage.Spanish, "Zagreb")
}
