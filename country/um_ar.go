//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsMinorOutlyingIslands.RegisterName(xlanguage.Arabic, "جزر الولايات المتحدة الصغيرة النائية")
	dataUsMinorOutlyingIslands.RegisterOfficialName(xlanguage.Arabic, "جزر الولايات المتحدة الصغيرة النائية")
	dataUsMinorOutlyingIslands.RegisterCapital(xlanguage.Arabic, "—")
}
