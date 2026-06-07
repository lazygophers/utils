package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMorocco.RegisterName(xlanguage.Arabic, "المغرب")
	dataMorocco.RegisterOfficialName(xlanguage.Arabic, "المملكة المغربية")
	dataMorocco.RegisterCapital(xlanguage.Arabic, "الرباط")
}
