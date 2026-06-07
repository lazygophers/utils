//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPhilippines.RegisterName(xlanguage.Spanish, "Filipinas")
	dataPhilippines.RegisterOfficialName(xlanguage.Spanish, "República de Filipinas")
	dataPhilippines.RegisterCapital(xlanguage.Spanish, "Manila")
}
