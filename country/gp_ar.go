//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_gp)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuadeloupe.RegisterName(xlanguage.Arabic, "غوادلوب")
	dataGuadeloupe.RegisterOfficialName(xlanguage.Arabic, "غوادلوب")
	dataGuadeloupe.RegisterCapital(xlanguage.Arabic, "باس تير")
}
