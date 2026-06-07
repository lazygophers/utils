package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataParaguay.RegisterName(xlanguage.Spanish, "Paraguay")
	dataParaguay.RegisterOfficialName(xlanguage.Spanish, "República del Paraguay")
	dataParaguay.RegisterCapital(xlanguage.Spanish, "Asunción")
}
