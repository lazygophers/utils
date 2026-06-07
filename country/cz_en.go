package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCzechia.RegisterName(xlanguage.English, "Czechia")
	dataCzechia.RegisterOfficialName(xlanguage.English, "Czech Republic")
	dataCzechia.RegisterCapital(xlanguage.English, "Prague")
}
