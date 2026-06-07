package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlbania.RegisterName(xlanguage.English, "Albania")
	dataAlbania.RegisterOfficialName(xlanguage.English, "Republic of Albania")
	dataAlbania.RegisterCapital(xlanguage.English, "Tirana")
}
