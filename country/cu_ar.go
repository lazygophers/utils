//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuba.RegisterName(xlanguage.Arabic, "كوبا")
	dataCuba.RegisterOfficialName(xlanguage.Arabic, "جمهورية كوبا")
	dataCuba.RegisterCapital(xlanguage.Arabic, "هافانا")
}
