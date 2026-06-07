package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuadeloupe.RegisterName(xlanguage.French, "Guadeloupe")
	dataGuadeloupe.RegisterOfficialName(xlanguage.French, "Guadeloupe")
	dataGuadeloupe.RegisterCapital(xlanguage.French, "Basse-Terre")
}
