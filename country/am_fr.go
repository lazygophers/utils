//go:build (lang_fr || lang_all) && (country_all || country_am || country_asia || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.French, "Arménie")
	dataArmenia.RegisterOfficialName(xlanguage.French, "République d'Arménie")
	dataArmenia.RegisterCapital(xlanguage.French, "Erevan")
}
