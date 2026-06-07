package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGermany.RegisterName(xlanguage.English, "Germany")
	dataGermany.RegisterOfficialName(xlanguage.English, "Federal Republic of Germany")
	dataGermany.RegisterCapital(xlanguage.English, "Berlin")
}
