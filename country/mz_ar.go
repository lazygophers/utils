//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_mz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMozambique.RegisterName(xlanguage.Arabic, "موزمبيق")
	dataMozambique.RegisterOfficialName(xlanguage.Arabic, "جمهورية موزمبيق")
	dataMozambique.RegisterCapital(xlanguage.Arabic, "مابوتو")
}
