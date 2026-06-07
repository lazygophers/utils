//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_cu)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuba.RegisterName(xlanguage.Arabic, "كوبا")
	dataCuba.RegisterOfficialName(xlanguage.Arabic, "جمهورية كوبا")
	dataCuba.RegisterCapital(xlanguage.Arabic, "هافانا")
}
