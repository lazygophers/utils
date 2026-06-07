package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHonduras.RegisterName(xlanguage.Spanish, "Honduras")
	dataHonduras.RegisterOfficialName(xlanguage.Spanish, "República de Honduras")
	dataHonduras.RegisterCapital(xlanguage.Spanish, "Tegucigalpa")
}
