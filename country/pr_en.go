package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPuertoRico.RegisterName(xlanguage.English, "Puerto Rico")
	dataPuertoRico.RegisterOfficialName(xlanguage.English, "Commonwealth of Puerto Rico")
	dataPuertoRico.RegisterCapital(xlanguage.English, "San Juan")
}
