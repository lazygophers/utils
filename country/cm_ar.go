//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.Arabic, "الكاميرون")
	dataCameroon.RegisterOfficialName(xlanguage.Arabic, "جمهورية الكاميرون")
	dataCameroon.RegisterCapital(xlanguage.Arabic, "ياوندي")
}
