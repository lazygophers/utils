package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFinland.RegisterName(xlanguage.English, "Finland")
	dataFinland.RegisterOfficialName(xlanguage.English, "Republic of Finland")
	dataFinland.RegisterCapital(xlanguage.English, "Helsinki")
}
