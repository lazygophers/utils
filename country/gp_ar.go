//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuadeloupe.RegisterName(xlanguage.Arabic, "غوادلوب")
	dataGuadeloupe.RegisterOfficialName(xlanguage.Arabic, "غوادلوب")
	dataGuadeloupe.RegisterCapital(xlanguage.Arabic, "باس تير")
}
