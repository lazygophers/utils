package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTanzania.RegisterName(xlanguage.English, "Tanzania")
	dataTanzania.RegisterOfficialName(xlanguage.English, "United Republic of Tanzania")
	dataTanzania.RegisterCapital(xlanguage.English, "Dodoma")
}
