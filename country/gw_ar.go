//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuineaBissau.RegisterName(xlanguage.Arabic, "غينيا بيساو")
	dataGuineaBissau.RegisterOfficialName(xlanguage.Arabic, "جمهورية غينيا بيساو")
	dataGuineaBissau.RegisterCapital(xlanguage.Arabic, "بيساو")
}
