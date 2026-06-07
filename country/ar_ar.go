//go:build (lang_ar || lang_all) && (country_all || country_americas || country_ar || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArgentina.RegisterName(xlanguage.Arabic, "الأرجنتين")
	dataArgentina.RegisterOfficialName(xlanguage.Arabic, "جمهورية الأرجنتين")
	dataArgentina.RegisterCapital(xlanguage.Arabic, "بوينس آيرس")
}
