package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBolivia.RegisterName(xlanguage.Spanish, "Bolivia")
	dataBolivia.RegisterOfficialName(xlanguage.Spanish, "Estado Plurinacional de Bolivia")
	dataBolivia.RegisterCapital(xlanguage.Spanish, "Sucre")
}
