//go:build (lang_ar || lang_all) && (country_all || country_europe || country_fi || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFinland.RegisterName(xlanguage.Arabic, "فنلندا")
	dataFinland.RegisterOfficialName(xlanguage.Arabic, "جمهورية فنلندا")
	dataFinland.RegisterCapital(xlanguage.Arabic, "هلسنكي")
}
