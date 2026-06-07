//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthSudan.RegisterName(xlanguage.Arabic, "جنوب السودان")
	dataSouthSudan.RegisterOfficialName(xlanguage.Arabic, "جمهورية جنوب السودان")
	dataSouthSudan.RegisterCapital(xlanguage.Arabic, "جوبا")
}
