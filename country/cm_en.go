package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.English, "Cameroon")
	dataCameroon.RegisterOfficialName(xlanguage.English, "Republic of Cameroon")
	dataCameroon.RegisterCapital(xlanguage.English, "Yaounde")
}
