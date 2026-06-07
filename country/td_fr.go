package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChad.RegisterName(xlanguage.French, "Tchad")
	dataChad.RegisterOfficialName(xlanguage.French, "République du Tchad")
	dataChad.RegisterCapital(xlanguage.French, "N'Djamena")
}
