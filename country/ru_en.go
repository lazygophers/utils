package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRussia.RegisterName(xlanguage.English, "Russia")
	dataRussia.RegisterOfficialName(xlanguage.English, "Russian Federation")
	dataRussia.RegisterCapital(xlanguage.English, "Moscow")
}
