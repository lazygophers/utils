//go:build (lang_ar || lang_all) && (country_africa || country_all || country_gw || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuineaBissau.RegisterName(xlanguage.Arabic, "غينيا بيساو")
	dataGuineaBissau.RegisterOfficialName(xlanguage.Arabic, "جمهورية غينيا بيساو")
	dataGuineaBissau.RegisterCapital(xlanguage.Arabic, "بيساو")
}
