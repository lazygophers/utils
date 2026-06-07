package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLibya.RegisterName(xlanguage.English, "Libya")
	dataLibya.RegisterOfficialName(xlanguage.English, "State of Libya")
	dataLibya.RegisterCapital(xlanguage.English, "Tripoli")
}
