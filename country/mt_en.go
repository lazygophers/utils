package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalta.RegisterName(xlanguage.English, "Malta")
	dataMalta.RegisterOfficialName(xlanguage.English, "Republic of Malta")
	dataMalta.RegisterCapital(xlanguage.English, "Valletta")
}
