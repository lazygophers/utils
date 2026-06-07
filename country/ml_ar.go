//go:build (lang_ar || lang_all) && (country_africa || country_all || country_ml || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMali.RegisterName(xlanguage.Arabic, "مالي")
	dataMali.RegisterOfficialName(xlanguage.Arabic, "جمهورية مالي")
	dataMali.RegisterCapital(xlanguage.Arabic, "باماكو")
}
