package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBolivia.RegisterName(xlanguage.English, "Bolivia")
	dataBolivia.RegisterOfficialName(xlanguage.English, "Plurinational State of Bolivia")
	dataBolivia.RegisterCapital(xlanguage.English, "Sucre")
}
