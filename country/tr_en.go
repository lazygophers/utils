package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkey.RegisterName(xlanguage.English, "Turkey")
	dataTurkey.RegisterOfficialName(xlanguage.English, "Republic of Turkiye")
	dataTurkey.RegisterCapital(xlanguage.English, "Ankara")
}
