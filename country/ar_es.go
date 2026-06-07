package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArgentina.RegisterName(xlanguage.Spanish, "Argentina")
	dataArgentina.RegisterOfficialName(xlanguage.Spanish, "República Argentina")
	dataArgentina.RegisterCapital(xlanguage.Spanish, "Buenos Aires")
}
