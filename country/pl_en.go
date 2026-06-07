package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPoland.RegisterName(xlanguage.English, "Poland")
	dataPoland.RegisterOfficialName(xlanguage.English, "Republic of Poland")
	dataPoland.RegisterCapital(xlanguage.English, "Warsaw")
}
