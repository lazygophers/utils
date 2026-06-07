//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.Russian, "Армения")
	dataArmenia.RegisterOfficialName(xlanguage.Russian, "Республика Армения")
	dataArmenia.RegisterCapital(xlanguage.Russian, "Ереван")
}
