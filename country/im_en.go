package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsleOfMan.RegisterName(xlanguage.English, "Isle of Man")
	dataIsleOfMan.RegisterOfficialName(xlanguage.English, "Isle of Man")
	dataIsleOfMan.RegisterCapital(xlanguage.English, "Douglas")
}
