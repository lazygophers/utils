//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArgentina.RegisterName(xlanguage.Arabic, "الأرجنتين")
	dataArgentina.RegisterOfficialName(xlanguage.Arabic, "جمهورية الأرجنتين")
	dataArgentina.RegisterCapital(xlanguage.Arabic, "بوينس آيرس")
}
