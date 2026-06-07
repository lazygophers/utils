package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNetherlands.RegisterName(xlanguage.English, "Netherlands")
	dataNetherlands.RegisterOfficialName(xlanguage.English, "Kingdom of the Netherlands")
	dataNetherlands.RegisterCapital(xlanguage.English, "Amsterdam")
}
