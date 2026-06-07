package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataColombia.RegisterName(xlanguage.English, "Colombia")
	dataColombia.RegisterOfficialName(xlanguage.English, "Republic of Colombia")
	dataColombia.RegisterCapital(xlanguage.English, "Bogota")
}
