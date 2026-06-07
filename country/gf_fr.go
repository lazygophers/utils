package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchGuiana.RegisterName(xlanguage.French, "Guyane")
	dataFrenchGuiana.RegisterOfficialName(xlanguage.French, "Guyane française")
	dataFrenchGuiana.RegisterCapital(xlanguage.French, "Cayenne")
}
