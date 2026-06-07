package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuinea.RegisterName(xlanguage.French, "Guinée")
	dataGuinea.RegisterOfficialName(xlanguage.French, "République de Guinée")
	dataGuinea.RegisterCapital(xlanguage.French, "Conakry")
}
