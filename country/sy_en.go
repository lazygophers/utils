package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSyria.RegisterName(xlanguage.English, "Syria")
	dataSyria.RegisterOfficialName(xlanguage.English, "Syrian Arab Republic")
	dataSyria.RegisterCapital(xlanguage.English, "Damascus")
}
