//go:build (lang_ar || lang_all) && (country_africa || country_all || country_gm || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGambia.RegisterName(xlanguage.Arabic, "غامبيا")
	dataGambia.RegisterOfficialName(xlanguage.Arabic, "جمهورية غامبيا")
	dataGambia.RegisterCapital(xlanguage.Arabic, "بانجول")
}
