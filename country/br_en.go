//go:build country_all || country_americas || country_br || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrazil.RegisterName(xlanguage.English, "Brazil")
	dataBrazil.RegisterOfficialName(xlanguage.English, "Federative Republic of Brazil")
	dataBrazil.RegisterCapital(xlanguage.English, "Brasilia")
}
