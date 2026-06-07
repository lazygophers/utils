package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIraq.RegisterName(xlanguage.English, "Iraq")
	dataIraq.RegisterOfficialName(xlanguage.English, "Republic of Iraq")
	dataIraq.RegisterCapital(xlanguage.English, "Baghdad")
}
