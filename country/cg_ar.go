//go:build (lang_ar || lang_all) && (country_africa || country_all || country_cg || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCongo.RegisterName(xlanguage.Arabic, "جمهورية الكونغو")
	dataCongo.RegisterOfficialName(xlanguage.Arabic, "جمهورية الكونغو")
	dataCongo.RegisterCapital(xlanguage.Arabic, "برازافيل")
}
