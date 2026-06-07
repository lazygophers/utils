//go:build (lang_ar || lang_all) && (country_all || country_asia || country_central_asia || country_tj)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTajikistan.RegisterName(xlanguage.Arabic, "طاجيكستان")
	dataTajikistan.RegisterOfficialName(xlanguage.Arabic, "جمهورية طاجيكستان")
	dataTajikistan.RegisterCapital(xlanguage.Arabic, "دوشنبه")
}
