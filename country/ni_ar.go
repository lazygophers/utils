//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNicaragua.RegisterName(xlanguage.Arabic, "نيكاراغوا")
	dataNicaragua.RegisterOfficialName(xlanguage.Arabic, "جمهورية نيكاراغوا")
	dataNicaragua.RegisterCapital(xlanguage.Arabic, "ماناغوا")
}
