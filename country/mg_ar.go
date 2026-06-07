//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_mg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMadagascar.RegisterName(xlanguage.Arabic, "مدغشقر")
	dataMadagascar.RegisterOfficialName(xlanguage.Arabic, "جمهورية مدغشقر")
	dataMadagascar.RegisterCapital(xlanguage.Arabic, "أنتاناناريفو")
}
