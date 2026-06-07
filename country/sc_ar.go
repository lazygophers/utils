//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_sc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSeychelles.RegisterName(xlanguage.Arabic, "سيشل")
	dataSeychelles.RegisterOfficialName(xlanguage.Arabic, "جمهورية سيشل")
	dataSeychelles.RegisterCapital(xlanguage.Arabic, "فيكتوريا")
}
