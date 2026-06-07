package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEstonia.RegisterName(xlanguage.English, "Estonia")
	dataEstonia.RegisterOfficialName(xlanguage.English, "Republic of Estonia")
	dataEstonia.RegisterCapital(xlanguage.English, "Tallinn")
}
