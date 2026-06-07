package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLatvia.RegisterName(xlanguage.English, "Latvia")
	dataLatvia.RegisterOfficialName(xlanguage.English, "Republic of Latvia")
	dataLatvia.RegisterCapital(xlanguage.English, "Riga")
}
