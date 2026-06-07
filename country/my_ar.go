//go:build (lang_ar || lang_all) && (country_all || country_asia || country_my || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalaysia.RegisterName(xlanguage.Arabic, "ماليزيا")
	dataMalaysia.RegisterOfficialName(xlanguage.Arabic, "ماليزيا")
	dataMalaysia.RegisterCapital(xlanguage.Arabic, "كوالالمبور")
}
