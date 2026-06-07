package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDrCongo.RegisterName(xlanguage.English, "DR Congo")
	dataDrCongo.RegisterOfficialName(xlanguage.English, "Democratic Republic of the Congo")
	dataDrCongo.RegisterCapital(xlanguage.English, "Kinshasa")
}
