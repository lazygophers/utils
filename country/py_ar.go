//go:build (lang_ar || lang_all) && (country_all || country_americas || country_py || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataParaguay.RegisterName(xlanguage.Arabic, "باراغواي")
	dataParaguay.RegisterOfficialName(xlanguage.Arabic, "جمهورية باراغواي")
	dataParaguay.RegisterCapital(xlanguage.Arabic, "أسونسيون")
}
