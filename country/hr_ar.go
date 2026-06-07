//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCroatia.RegisterName(xlanguage.Arabic, "كرواتيا")
	dataCroatia.RegisterOfficialName(xlanguage.Arabic, "جمهورية كرواتيا")
	dataCroatia.RegisterCapital(xlanguage.Arabic, "زغرب")
}
