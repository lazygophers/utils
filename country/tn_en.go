package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTunisia.RegisterName(xlanguage.English, "Tunisia")
	dataTunisia.RegisterOfficialName(xlanguage.English, "Republic of Tunisia")
	dataTunisia.RegisterCapital(xlanguage.English, "Tunis")
}
