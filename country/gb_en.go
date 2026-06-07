package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedKingdom.RegisterName(xlanguage.English, "United Kingdom")
	dataUnitedKingdom.RegisterOfficialName(xlanguage.English, "United Kingdom of Great Britain and Northern Ireland")
	dataUnitedKingdom.RegisterCapital(xlanguage.English, "London")
}
