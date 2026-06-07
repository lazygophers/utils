//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_mw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalawi.RegisterName(xlanguage.Arabic, "مالاوي")
	dataMalawi.RegisterOfficialName(xlanguage.Arabic, "جمهورية مالاوي")
	dataMalawi.RegisterCapital(xlanguage.Arabic, "ليلونغوي")
}
