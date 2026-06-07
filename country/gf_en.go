package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchGuiana.RegisterName(xlanguage.English, "French Guiana")
	dataFrenchGuiana.RegisterOfficialName(xlanguage.English, "Guiana")
	dataFrenchGuiana.RegisterCapital(xlanguage.English, "Cayenne")
}
