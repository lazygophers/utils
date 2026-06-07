//go:build (lang_ar || lang_all) && (country_all || country_gu || country_micronesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuam.RegisterName(xlanguage.Arabic, "غوام")
	dataGuam.RegisterOfficialName(xlanguage.Arabic, "إقليم غوام")
	dataGuam.RegisterCapital(xlanguage.Arabic, "هاغاتنيا")
}
