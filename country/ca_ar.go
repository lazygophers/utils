//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCanada.RegisterName(xlanguage.Arabic, "كندا")
	dataCanada.RegisterOfficialName(xlanguage.Arabic, "كندا")
	dataCanada.RegisterCapital(xlanguage.Arabic, "أوتاوا")
}
