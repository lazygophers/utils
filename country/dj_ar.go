//go:build country_africa || country_all || country_dj || country_eastern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDjibouti.RegisterName(xlanguage.Arabic, "جيبوتي")
	dataDjibouti.RegisterOfficialName(xlanguage.Arabic, "جمهورية جيبوتي")
	dataDjibouti.RegisterCapital(xlanguage.Arabic, "مدينة جيبوتي")
}
