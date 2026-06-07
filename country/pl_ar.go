//go:build (lang_ar || lang_all) && (country_all || country_eastern_europe || country_europe || country_pl)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPoland.RegisterName(xlanguage.Arabic, "بولندا")
	dataPoland.RegisterOfficialName(xlanguage.Arabic, "جمهورية بولندا")
	dataPoland.RegisterCapital(xlanguage.Arabic, "وارسو")
}
