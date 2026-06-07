package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCroatia.RegisterName(xlanguage.English, "Croatia")
	dataCroatia.RegisterOfficialName(xlanguage.English, "Republic of Croatia")
	dataCroatia.RegisterCapital(xlanguage.English, "Zagreb")
}
