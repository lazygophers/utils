package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAngola.RegisterName(xlanguage.English, "Angola")
	dataAngola.RegisterOfficialName(xlanguage.English, "Republic of Angola")
	dataAngola.RegisterCapital(xlanguage.English, "Luanda")
}
