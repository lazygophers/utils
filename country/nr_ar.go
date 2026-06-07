//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNauru.RegisterName(xlanguage.Arabic, "ناورو")
	dataNauru.RegisterOfficialName(xlanguage.Arabic, "جمهورية ناورو")
	dataNauru.RegisterCapital(xlanguage.Arabic, "ياريل")
}
