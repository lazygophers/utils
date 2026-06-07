package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIreland.RegisterName(xlanguage.English, "Ireland")
	dataIreland.RegisterOfficialName(xlanguage.English, "Republic of Ireland")
	dataIreland.RegisterCapital(xlanguage.English, "Dublin")
}
