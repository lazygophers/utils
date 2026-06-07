package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTogo.RegisterName(xlanguage.English, "Togo")
	dataTogo.RegisterOfficialName(xlanguage.English, "Togolese Republic")
	dataTogo.RegisterCapital(xlanguage.English, "Lome")
}
