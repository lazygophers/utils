//go:build (lang_ar || lang_all) && (country_all || country_micronesia || country_oceania || country_um)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsMinorOutlyingIslands.RegisterName(xlanguage.Arabic, "جزر الولايات المتحدة الصغيرة النائية")
	dataUsMinorOutlyingIslands.RegisterOfficialName(xlanguage.Arabic, "جزر الولايات المتحدة الصغيرة النائية")
	dataUsMinorOutlyingIslands.RegisterCapital(xlanguage.Arabic, "—")
}
