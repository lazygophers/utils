package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMaldives.RegisterName(xlanguage.English, "Maldives")
	dataMaldives.RegisterOfficialName(xlanguage.English, "Republic of Maldives")
	dataMaldives.RegisterCapital(xlanguage.English, "Male")
}
