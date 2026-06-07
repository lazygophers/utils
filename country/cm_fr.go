package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.French, "Cameroun")
	dataCameroon.RegisterOfficialName(xlanguage.French, "République du Cameroun")
	dataCameroon.RegisterCapital(xlanguage.French, "Yaoundé")
}
