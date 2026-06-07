package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUzbekistan.RegisterName(xlanguage.English, "Uzbekistan")
	dataUzbekistan.RegisterOfficialName(xlanguage.English, "Republic of Uzbekistan")
	dataUzbekistan.RegisterCapital(xlanguage.English, "Tashkent")
}
