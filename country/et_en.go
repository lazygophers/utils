package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEthiopia.RegisterName(xlanguage.English, "Ethiopia")
	dataEthiopia.RegisterOfficialName(xlanguage.English, "Federal Democratic Republic of Ethiopia")
	dataEthiopia.RegisterCapital(xlanguage.English, "Addis Ababa")
}
