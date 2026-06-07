package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsrael.RegisterName(xlanguage.English, "Israel")
	dataIsrael.RegisterOfficialName(xlanguage.English, "State of Israel")
	dataIsrael.RegisterCapital(xlanguage.English, "Jerusalem")
}
