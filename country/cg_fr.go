package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCongo.RegisterName(xlanguage.French, "République du Congo")
	dataCongo.RegisterOfficialName(xlanguage.French, "République du Congo")
	dataCongo.RegisterCapital(xlanguage.French, "Brazzaville")
}
