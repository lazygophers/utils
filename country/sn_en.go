package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSenegal.RegisterName(xlanguage.English, "Senegal")
	dataSenegal.RegisterOfficialName(xlanguage.English, "Republic of Senegal")
	dataSenegal.RegisterCapital(xlanguage.English, "Dakar")
}
