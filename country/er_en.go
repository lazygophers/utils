package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEritrea.RegisterName(xlanguage.English, "Eritrea")
	dataEritrea.RegisterOfficialName(xlanguage.English, "State of Eritrea")
	dataEritrea.RegisterCapital(xlanguage.English, "Asmara")
}
