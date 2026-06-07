//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedKingdom.RegisterName(xlanguage.Spanish, "Reino Unido")
	dataUnitedKingdom.RegisterOfficialName(xlanguage.Spanish, "Reino Unido de Gran Bretaña e Irlanda del Norte")
	dataUnitedKingdom.RegisterCapital(xlanguage.Spanish, "Londres")
}
