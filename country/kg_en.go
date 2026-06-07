package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKyrgyzstan.RegisterName(xlanguage.English, "Kyrgyzstan")
	dataKyrgyzstan.RegisterOfficialName(xlanguage.English, "Kyrgyz Republic")
	dataKyrgyzstan.RegisterCapital(xlanguage.English, "Bishkek")
}
