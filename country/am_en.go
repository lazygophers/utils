package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.English, "Armenia")
	dataArmenia.RegisterOfficialName(xlanguage.English, "Republic of Armenia")
	dataArmenia.RegisterCapital(xlanguage.English, "Yerevan")
}
