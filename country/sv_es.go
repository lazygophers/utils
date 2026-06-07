package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataElSalvador.RegisterName(xlanguage.Spanish, "El Salvador")
	dataElSalvador.RegisterOfficialName(xlanguage.Spanish, "República de El Salvador")
	dataElSalvador.RegisterCapital(xlanguage.Spanish, "San Salvador")
}
