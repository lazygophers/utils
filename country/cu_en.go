package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuba.RegisterName(xlanguage.English, "Cuba")
	dataCuba.RegisterOfficialName(xlanguage.English, "Republic of Cuba")
	dataCuba.RegisterCapital(xlanguage.English, "Havana")
}
