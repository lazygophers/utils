//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUganda.RegisterName(xlanguage.Arabic, "أوغندا")
	dataUganda.RegisterOfficialName(xlanguage.Arabic, "جمهورية أوغندا")
	dataUganda.RegisterCapital(xlanguage.Arabic, "كمبالا")
}
