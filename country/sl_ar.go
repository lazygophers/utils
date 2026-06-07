//go:build (lang_ar || lang_all) && (country_africa || country_all || country_sl || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSierraLeone.RegisterName(xlanguage.Arabic, "سيراليون")
	dataSierraLeone.RegisterOfficialName(xlanguage.Arabic, "جمهورية سيراليون")
	dataSierraLeone.RegisterCapital(xlanguage.Arabic, "فريتاون")
}
