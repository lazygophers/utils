//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntiguaAndBarbuda.RegisterName(xlanguage.Arabic, "أنتيغوا وباربودا")
	dataAntiguaAndBarbuda.RegisterOfficialName(xlanguage.Arabic, "أنتيغوا وباربودا")
	dataAntiguaAndBarbuda.RegisterCapital(xlanguage.Arabic, "سانت جونز")
}
