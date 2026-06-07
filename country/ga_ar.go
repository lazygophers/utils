//go:build (lang_ar || lang_all) && (country_africa || country_all || country_ga || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGabon.RegisterName(xlanguage.Arabic, "الغابون")
	dataGabon.RegisterOfficialName(xlanguage.Arabic, "الجمهورية الغابونية")
	dataGabon.RegisterCapital(xlanguage.Arabic, "ليبرفيل")
}
