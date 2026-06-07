//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_gd)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGrenada.RegisterName(xlanguage.Arabic, "غرينادا")
	dataGrenada.RegisterOfficialName(xlanguage.Arabic, "غرينادا")
	dataGrenada.RegisterCapital(xlanguage.Arabic, "سانت جورج")
}
