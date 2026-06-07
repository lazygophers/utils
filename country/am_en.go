//go:build country_all || country_am || country_asia || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.English, "Armenia")
	dataArmenia.RegisterOfficialName(xlanguage.English, "Republic of Armenia")
	dataArmenia.RegisterCapital(xlanguage.English, "Yerevan")
}
