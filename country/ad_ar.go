//go:build (lang_ar || lang_all) && (country_ad || country_all || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAndorra.RegisterName(xlanguage.Arabic, "أندورا")
	dataAndorra.RegisterOfficialName(xlanguage.Arabic, "إمارة أندورا")
	dataAndorra.RegisterCapital(xlanguage.Arabic, "أندورا لا فيلا")
}
