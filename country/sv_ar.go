//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataElSalvador.RegisterName(xlanguage.Arabic, "السلفادور")
	dataElSalvador.RegisterOfficialName(xlanguage.Arabic, "جمهورية السلفادور")
	dataElSalvador.RegisterCapital(xlanguage.Arabic, "سان سلفادور")
}
