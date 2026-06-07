package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHaiti.RegisterName(xlanguage.French, "Haïti")
	dataHaiti.RegisterOfficialName(xlanguage.French, "République d'Haïti")
	dataHaiti.RegisterCapital(xlanguage.French, "Port-au-Prince")
}
