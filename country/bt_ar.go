//go:build (lang_ar || lang_all) && (country_all || country_asia || country_bt || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBhutan.RegisterName(xlanguage.Arabic, "بوتان")
	dataBhutan.RegisterOfficialName(xlanguage.Arabic, "مملكة بوتان")
	dataBhutan.RegisterCapital(xlanguage.Arabic, "تيمفو")
}
