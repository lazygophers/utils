package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEcuador.RegisterName(xlanguage.Spanish, "Ecuador")
	dataEcuador.RegisterOfficialName(xlanguage.Spanish, "República del Ecuador")
	dataEcuador.RegisterCapital(xlanguage.Spanish, "Quito")
}
