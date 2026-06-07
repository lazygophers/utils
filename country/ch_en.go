package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSwitzerland.RegisterName(xlanguage.English, "Switzerland")
	dataSwitzerland.RegisterOfficialName(xlanguage.English, "Swiss Confederation")
	dataSwitzerland.RegisterCapital(xlanguage.English, "Bern")
}
