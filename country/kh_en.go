package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.English, "Cambodia")
	dataCambodia.RegisterOfficialName(xlanguage.English, "Kingdom of Cambodia")
	dataCambodia.RegisterCapital(xlanguage.English, "Phnom Penh")
}
