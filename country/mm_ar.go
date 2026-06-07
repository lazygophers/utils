//go:build (lang_ar || lang_all) && (country_all || country_asia || country_mm || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMyanmar.RegisterName(xlanguage.Arabic, "ميانمار")
	dataMyanmar.RegisterOfficialName(xlanguage.Arabic, "جمهورية اتحاد ميانمار")
	dataMyanmar.RegisterCapital(xlanguage.Arabic, "نايبيداو")
}
