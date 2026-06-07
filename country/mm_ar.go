//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMyanmar.RegisterName(xlanguage.Arabic, "ميانمار")
	dataMyanmar.RegisterOfficialName(xlanguage.Arabic, "جمهورية اتحاد ميانمار")
	dataMyanmar.RegisterCapital(xlanguage.Arabic, "نايبيداو")
}
