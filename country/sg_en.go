package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSingapore.RegisterName(xlanguage.English, "Singapore")
	dataSingapore.RegisterOfficialName(xlanguage.English, "Republic of Singapore")
	dataSingapore.RegisterCapital(xlanguage.English, "Singapore")
}
