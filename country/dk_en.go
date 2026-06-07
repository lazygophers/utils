package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDenmark.RegisterName(xlanguage.English, "Denmark")
	dataDenmark.RegisterOfficialName(xlanguage.English, "Kingdom of Denmark")
	dataDenmark.RegisterCapital(xlanguage.English, "Copenhagen")
}
