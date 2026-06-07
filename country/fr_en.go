package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrance.RegisterName(xlanguage.English, "France")
	dataFrance.RegisterOfficialName(xlanguage.English, "French Republic")
	dataFrance.RegisterCapital(xlanguage.English, "Paris")
}
