package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUkraine.RegisterName(xlanguage.English, "Ukraine")
	dataUkraine.RegisterOfficialName(xlanguage.English, "Ukraine")
	dataUkraine.RegisterCapital(xlanguage.English, "Kyiv")
}
