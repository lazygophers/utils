package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrazil.RegisterName(xlanguage.English, "Brazil")
	dataBrazil.RegisterOfficialName(xlanguage.English, "Federative Republic of Brazil")
	dataBrazil.RegisterCapital(xlanguage.English, "Brasilia")
}
