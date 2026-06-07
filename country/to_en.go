package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTonga.RegisterName(xlanguage.English, "Tonga")
	dataTonga.RegisterOfficialName(xlanguage.English, "Kingdom of Tonga")
	dataTonga.RegisterCapital(xlanguage.English, "Nuku'alofa")
}
