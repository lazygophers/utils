//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMyanmar.RegisterName(xlanguage.Spanish, "Birmania")
	dataMyanmar.RegisterOfficialName(xlanguage.Spanish, "República de la Unión de Myanmar")
	dataMyanmar.RegisterCapital(xlanguage.Spanish, "Naipyidó")
}
