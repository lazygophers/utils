package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuyana.RegisterName(xlanguage.English, "Guyana")
	dataGuyana.RegisterOfficialName(xlanguage.English, "Co-operative Republic of Guyana")
	dataGuyana.RegisterCapital(xlanguage.English, "Georgetown")
}
