package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJapan.RegisterName(xlanguage.English, "Japan")
	dataJapan.RegisterOfficialName(xlanguage.English, "Japan")
	dataJapan.RegisterCapital(xlanguage.English, "Tokyo")
}
