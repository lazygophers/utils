//go:build (lang_ar || lang_all) && (country_ag || country_all || country_americas || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntiguaAndBarbuda.RegisterName(xlanguage.Arabic, "أنتيغوا وباربودا")
	dataAntiguaAndBarbuda.RegisterOfficialName(xlanguage.Arabic, "أنتيغوا وباربودا")
	dataAntiguaAndBarbuda.RegisterCapital(xlanguage.Arabic, "سانت جونز")
}
