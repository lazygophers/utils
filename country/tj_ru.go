package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTajikistan.RegisterName(xlanguage.Russian, "Таджикистан")
	dataTajikistan.RegisterOfficialName(xlanguage.Russian, "Республика Таджикистан")
	dataTajikistan.RegisterCapital(xlanguage.Russian, "Душанбе")
}
