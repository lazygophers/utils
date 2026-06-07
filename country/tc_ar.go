//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurksAndCaicosIslands.RegisterName(xlanguage.Arabic, "جزر توركس وكايكوس")
	dataTurksAndCaicosIslands.RegisterOfficialName(xlanguage.Arabic, "إقليم جزر توركس وكايكوس البريطاني فيما وراء البحار")
	dataTurksAndCaicosIslands.RegisterCapital(xlanguage.Arabic, "كوكبيرن تاون")
}
