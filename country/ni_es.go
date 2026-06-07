package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNicaragua.RegisterName(xlanguage.Spanish, "Nicaragua")
	dataNicaragua.RegisterOfficialName(xlanguage.Spanish, "República de Nicaragua")
	dataNicaragua.RegisterCapital(xlanguage.Spanish, "Managua")
}
