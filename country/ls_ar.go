//go:build (lang_ar || lang_all) && (country_africa || country_all || country_ls || country_southern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLesotho.RegisterName(xlanguage.Arabic, "ليسوتو")
	dataLesotho.RegisterOfficialName(xlanguage.Arabic, "مملكة ليسوتو")
	dataLesotho.RegisterCapital(xlanguage.Arabic, "ماسيرو")
}
