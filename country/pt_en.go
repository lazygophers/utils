package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPortugal.RegisterName(xlanguage.English, "Portugal")
	dataPortugal.RegisterOfficialName(xlanguage.English, "Portuguese Republic")
	dataPortugal.RegisterCapital(xlanguage.English, "Lisbon")
}
