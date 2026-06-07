package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrance.RegisterName(xlanguage.French, "France")
	dataFrance.RegisterOfficialName(xlanguage.French, "République française")
	dataFrance.RegisterCapital(xlanguage.French, "Paris")
}
