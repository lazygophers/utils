//go:build (lang_ru || lang_all) && (country_all || country_am || country_asia || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.Russian, "Армения")
	dataArmenia.RegisterOfficialName(xlanguage.Russian, "Республика Армения")
	dataArmenia.RegisterCapital(xlanguage.Russian, "Ереван")
}
