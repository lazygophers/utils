//go:build country_all || country_americas || country_central_america || country_sv

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataElSalvador.RegisterName(xlanguage.Spanish, "El Salvador")
	dataElSalvador.RegisterOfficialName(xlanguage.Spanish, "República de El Salvador")
	dataElSalvador.RegisterCapital(xlanguage.Spanish, "San Salvador")
}
