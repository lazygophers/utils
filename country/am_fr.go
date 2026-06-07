//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.French, "Arménie")
	dataArmenia.RegisterOfficialName(xlanguage.French, "République d'Arménie")
	dataArmenia.RegisterCapital(xlanguage.French, "Erevan")
}
