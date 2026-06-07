//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRwanda.RegisterName(xlanguage.Arabic, "رواندا")
	dataRwanda.RegisterOfficialName(xlanguage.Arabic, "جمهورية رواندا")
	dataRwanda.RegisterCapital(xlanguage.Arabic, "كيغالي")
}
