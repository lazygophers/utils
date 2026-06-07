package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJordan.RegisterName(xlanguage.English, "Jordan")
	dataJordan.RegisterOfficialName(xlanguage.English, "Hashemite Kingdom of Jordan")
	dataJordan.RegisterCapital(xlanguage.English, "Amman")
}
