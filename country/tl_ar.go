//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTimorLeste.RegisterName(xlanguage.Arabic, "تيمور الشرقية")
	dataTimorLeste.RegisterOfficialName(xlanguage.Arabic, "جمهورية تيمور الشرقية الديمقراطية")
	dataTimorLeste.RegisterCapital(xlanguage.Arabic, "ديلي")
}
