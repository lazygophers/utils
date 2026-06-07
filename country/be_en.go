package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelgium.RegisterName(xlanguage.English, "Belgium")
	dataBelgium.RegisterOfficialName(xlanguage.English, "Kingdom of Belgium")
	dataBelgium.RegisterCapital(xlanguage.English, "Brussels")
}
