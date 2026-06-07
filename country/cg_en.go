package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCongo.RegisterName(xlanguage.English, "Congo")
	dataCongo.RegisterOfficialName(xlanguage.English, "Republic of the Congo")
	dataCongo.RegisterCapital(xlanguage.English, "Brazzaville")
}
