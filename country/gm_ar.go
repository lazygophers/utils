//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGambia.RegisterName(xlanguage.Arabic, "غامبيا")
	dataGambia.RegisterOfficialName(xlanguage.Arabic, "جمهورية غامبيا")
	dataGambia.RegisterCapital(xlanguage.Arabic, "بانجول")
}
