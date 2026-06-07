package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuba.RegisterName(xlanguage.Spanish, "Cuba")
	dataCuba.RegisterOfficialName(xlanguage.Spanish, "República de Cuba")
	dataCuba.RegisterCapital(xlanguage.Spanish, "La Habana")
}
