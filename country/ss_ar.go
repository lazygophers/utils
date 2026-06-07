//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_ss)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthSudan.RegisterName(xlanguage.Arabic, "جنوب السودان")
	dataSouthSudan.RegisterOfficialName(xlanguage.Arabic, "جمهورية جنوب السودان")
	dataSouthSudan.RegisterCapital(xlanguage.Arabic, "جوبا")
}
