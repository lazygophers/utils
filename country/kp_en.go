package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthKorea.RegisterName(xlanguage.English, "North Korea")
	dataNorthKorea.RegisterOfficialName(xlanguage.English, "Democratic People's Republic of Korea")
	dataNorthKorea.RegisterCapital(xlanguage.English, "Pyongyang")
}
