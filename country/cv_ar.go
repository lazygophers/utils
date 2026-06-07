//go:build (lang_ar || lang_all) && (country_africa || country_all || country_cv || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaboVerde.RegisterName(xlanguage.Arabic, "الرأس الأخضر")
	dataCaboVerde.RegisterOfficialName(xlanguage.Arabic, "جمهورية الرأس الأخضر")
	dataCaboVerde.RegisterCapital(xlanguage.Arabic, "برايا")
}
