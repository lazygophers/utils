package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGabon.RegisterName(xlanguage.French, "Gabon")
	dataGabon.RegisterOfficialName(xlanguage.French, "République gabonaise")
	dataGabon.RegisterCapital(xlanguage.French, "Libreville")
}
