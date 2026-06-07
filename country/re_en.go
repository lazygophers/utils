package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataReunion.RegisterName(xlanguage.English, "Reunion")
	dataReunion.RegisterOfficialName(xlanguage.English, "Reunion Island")
	dataReunion.RegisterCapital(xlanguage.English, "Saint-Denis")
}
