//go:build country_all || country_americas || country_central_america || country_sv

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataElSalvador.RegisterName(xlanguage.English, "El Salvador")
	dataElSalvador.RegisterOfficialName(xlanguage.English, "Republic of El Salvador")
	dataElSalvador.RegisterCapital(xlanguage.English, "San Salvador")
}
