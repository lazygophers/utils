package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovenia.RegisterName(xlanguage.English, "Slovenia")
	dataSlovenia.RegisterOfficialName(xlanguage.English, "Republic of Slovenia")
	dataSlovenia.RegisterCapital(xlanguage.English, "Ljubljana")
}
