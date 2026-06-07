package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChile.RegisterName(xlanguage.English, "Chile")
	dataChile.RegisterOfficialName(xlanguage.English, "Republic of Chile")
	dataChile.RegisterCapital(xlanguage.English, "Santiago")
}
