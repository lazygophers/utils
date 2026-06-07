package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataReunion.RegisterName(xlanguage.French, "La Réunion")
	dataReunion.RegisterOfficialName(xlanguage.French, "La Réunion")
	dataReunion.RegisterCapital(xlanguage.French, "Saint-Denis")
}
