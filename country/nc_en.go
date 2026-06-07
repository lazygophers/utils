package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewCaledonia.RegisterName(xlanguage.English, "New Caledonia")
	dataNewCaledonia.RegisterOfficialName(xlanguage.English, "New Caledonia")
	dataNewCaledonia.RegisterCapital(xlanguage.English, "Noumea")
}
