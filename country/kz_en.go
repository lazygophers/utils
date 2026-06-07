package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKazakhstan.RegisterName(xlanguage.English, "Kazakhstan")
	dataKazakhstan.RegisterOfficialName(xlanguage.English, "Republic of Kazakhstan")
	dataKazakhstan.RegisterCapital(xlanguage.English, "Astana")
}
