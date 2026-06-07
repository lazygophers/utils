//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVenezuela.RegisterName(xlanguage.Arabic, "فنزويلا")
	dataVenezuela.RegisterOfficialName(xlanguage.Arabic, "جمهورية فنزويلا البوليفارية")
	dataVenezuela.RegisterCapital(xlanguage.Arabic, "كاراكاس")
}
