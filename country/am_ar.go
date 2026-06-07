//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.Arabic, "أرمينيا")
	dataArmenia.RegisterOfficialName(xlanguage.Arabic, "جمهورية أرمينيا")
	dataArmenia.RegisterCapital(xlanguage.Arabic, "يريفان")
}
