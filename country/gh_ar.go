//go:build (lang_ar || lang_all) && (country_africa || country_all || country_gh || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGhana.RegisterName(xlanguage.Arabic, "غانا")
	dataGhana.RegisterOfficialName(xlanguage.Arabic, "جمهورية غانا")
	dataGhana.RegisterCapital(xlanguage.Arabic, "أكرا")
}
