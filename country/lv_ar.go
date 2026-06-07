//go:build (lang_ar || lang_all) && (country_all || country_europe || country_lv || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLatvia.RegisterName(xlanguage.Arabic, "لاتفيا")
	dataLatvia.RegisterOfficialName(xlanguage.Arabic, "جمهورية لاتفيا")
	dataLatvia.RegisterCapital(xlanguage.Arabic, "ريغا")
}
