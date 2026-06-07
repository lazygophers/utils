//go:build (lang_ar || lang_all) && (country_all || country_nu || country_oceania || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiue.RegisterName(xlanguage.Arabic, "نييوي")
	dataNiue.RegisterOfficialName(xlanguage.Arabic, "نييوي")
	dataNiue.RegisterCapital(xlanguage.Arabic, "ألوفي")
}
