package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBulgaria.RegisterName(xlanguage.English, "Bulgaria")
	dataBulgaria.RegisterOfficialName(xlanguage.English, "Republic of Bulgaria")
	dataBulgaria.RegisterCapital(xlanguage.English, "Sofia")
}
