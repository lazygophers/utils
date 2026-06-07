package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCanada.RegisterName(xlanguage.English, "Canada")
	dataCanada.RegisterOfficialName(xlanguage.English, "Canada")
	dataCanada.RegisterCapital(xlanguage.English, "Ottawa")
}
