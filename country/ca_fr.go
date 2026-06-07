package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCanada.RegisterName(xlanguage.French, "Canada")
	dataCanada.RegisterOfficialName(xlanguage.French, "Canada")
	dataCanada.RegisterCapital(xlanguage.French, "Ottawa")
}
