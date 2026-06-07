//go:build (lang_ar || lang_all) && (country_africa || country_all || country_ci || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIvoryCoast.RegisterName(xlanguage.Arabic, "ساحل العاج")
	dataIvoryCoast.RegisterOfficialName(xlanguage.Arabic, "جمهورية ساحل العاج")
	dataIvoryCoast.RegisterCapital(xlanguage.Arabic, "ياموسوكرو")
}
