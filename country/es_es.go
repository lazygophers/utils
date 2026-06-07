package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSpain.RegisterName(xlanguage.Spanish, "España")
	dataSpain.RegisterOfficialName(xlanguage.Spanish, "Reino de España")
	dataSpain.RegisterCapital(xlanguage.Spanish, "Madrid")
}
