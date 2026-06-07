//go:build (lang_ar || lang_all) && (country_all || country_americas || country_co || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataColombia.RegisterName(xlanguage.Arabic, "كولومبيا")
	dataColombia.RegisterOfficialName(xlanguage.Arabic, "جمهورية كولومبيا")
	dataColombia.RegisterCapital(xlanguage.Arabic, "بوغوتا")
}
