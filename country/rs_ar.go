//go:build (lang_ar || lang_all) && (country_all || country_europe || country_rs || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSerbia.RegisterName(xlanguage.Arabic, "صربيا")
	dataSerbia.RegisterOfficialName(xlanguage.Arabic, "جمهورية صربيا")
	dataSerbia.RegisterCapital(xlanguage.Arabic, "بلغراد")
}
