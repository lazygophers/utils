package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSudan.RegisterName(xlanguage.English, "Sudan")
	dataSudan.RegisterOfficialName(xlanguage.English, "Republic of the Sudan")
	dataSudan.RegisterCapital(xlanguage.English, "Khartoum")
}
