//go:build (lang_ar || lang_all) && (country_all || country_asia || country_kh || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.Arabic, "كمبوديا")
	dataCambodia.RegisterOfficialName(xlanguage.Arabic, "مملكة كمبوديا")
	dataCambodia.RegisterCapital(xlanguage.Arabic, "بنوم بنه")
}
