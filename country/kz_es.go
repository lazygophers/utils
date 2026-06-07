//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKazakhstan.RegisterName(xlanguage.Spanish, "Kazajistán")
	dataKazakhstan.RegisterOfficialName(xlanguage.Spanish, "República de Kazajistán")
	dataKazakhstan.RegisterCapital(xlanguage.Spanish, "Astaná")
}
