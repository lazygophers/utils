package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMoldova.RegisterName(xlanguage.English, "Moldova")
	dataMoldova.RegisterOfficialName(xlanguage.English, "Republic of Moldova")
	dataMoldova.RegisterCapital(xlanguage.English, "Chisinau")
}
