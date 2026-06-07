package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAzerbaijan.RegisterName(xlanguage.English, "Azerbaijan")
	dataAzerbaijan.RegisterOfficialName(xlanguage.English, "Republic of Azerbaijan")
	dataAzerbaijan.RegisterCapital(xlanguage.English, "Baku")
}
