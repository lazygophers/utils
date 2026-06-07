package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahamas.RegisterName(xlanguage.English, "Bahamas")
	dataBahamas.RegisterOfficialName(xlanguage.English, "Commonwealth of the Bahamas")
	dataBahamas.RegisterCapital(xlanguage.English, "Nassau")
}
