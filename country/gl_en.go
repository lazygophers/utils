package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreenland.RegisterName(xlanguage.English, "Greenland")
	dataGreenland.RegisterOfficialName(xlanguage.English, "Greenland")
	dataGreenland.RegisterCapital(xlanguage.English, "Nuuk")
}
