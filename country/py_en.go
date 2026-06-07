package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataParaguay.RegisterName(xlanguage.English, "Paraguay")
	dataParaguay.RegisterOfficialName(xlanguage.English, "Republic of Paraguay")
	dataParaguay.RegisterCapital(xlanguage.English, "Asuncion")
}
