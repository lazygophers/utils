package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChile.RegisterName(xlanguage.Spanish, "Chile")
	dataChile.RegisterOfficialName(xlanguage.Spanish, "República de Chile")
	dataChile.RegisterCapital(xlanguage.Spanish, "Santiago")
}
