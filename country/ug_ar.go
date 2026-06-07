//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_ug)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUganda.RegisterName(xlanguage.Arabic, "أوغندا")
	dataUganda.RegisterOfficialName(xlanguage.Arabic, "جمهورية أوغندا")
	dataUganda.RegisterCapital(xlanguage.Arabic, "كمبالا")
}
