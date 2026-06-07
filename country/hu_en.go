package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHungary.RegisterName(xlanguage.English, "Hungary")
	dataHungary.RegisterOfficialName(xlanguage.English, "Hungary")
	dataHungary.RegisterCapital(xlanguage.English, "Budapest")
}
