package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthKorea.RegisterName(xlanguage.English, "South Korea")
	dataSouthKorea.RegisterOfficialName(xlanguage.English, "Republic of Korea")
	dataSouthKorea.RegisterCapital(xlanguage.English, "Seoul")
}
