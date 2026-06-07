//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkey.RegisterName(xlanguage.Arabic, "تركيا")
	dataTurkey.RegisterOfficialName(xlanguage.Arabic, "جمهورية تركيا")
	dataTurkey.RegisterCapital(xlanguage.Arabic, "أنقرة")
}
