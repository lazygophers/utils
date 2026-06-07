//go:build (lang_ar || lang_all) && (country_all || country_americas || country_ca || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCanada.RegisterName(xlanguage.Arabic, "كندا")
	dataCanada.RegisterOfficialName(xlanguage.Arabic, "كندا")
	dataCanada.RegisterCapital(xlanguage.Arabic, "أوتاوا")
}
