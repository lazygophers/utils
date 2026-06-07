//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHaiti.RegisterName(xlanguage.Arabic, "هايتي")
	dataHaiti.RegisterOfficialName(xlanguage.Arabic, "جمهورية هايتي")
	dataHaiti.RegisterCapital(xlanguage.Arabic, "بورت أو برنس")
}
