package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataItaly.RegisterName(xlanguage.English, "Italy")
	dataItaly.RegisterOfficialName(xlanguage.English, "Italian Republic")
	dataItaly.RegisterCapital(xlanguage.English, "Rome")
}
