package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataThailand.RegisterName(xlanguage.English, "Thailand")
	dataThailand.RegisterOfficialName(xlanguage.English, "Kingdom of Thailand")
	dataThailand.RegisterCapital(xlanguage.English, "Bangkok")
}
