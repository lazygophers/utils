//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGermany.RegisterName(xlanguage.Arabic, "ألمانيا")
	dataGermany.RegisterOfficialName(xlanguage.Arabic, "جمهورية ألمانيا الاتحادية")
	dataGermany.RegisterCapital(xlanguage.Arabic, "برلين")
}
