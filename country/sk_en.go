package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovakia.RegisterName(xlanguage.English, "Slovakia")
	dataSlovakia.RegisterOfficialName(xlanguage.English, "Slovak Republic")
	dataSlovakia.RegisterCapital(xlanguage.English, "Bratislava")
}
