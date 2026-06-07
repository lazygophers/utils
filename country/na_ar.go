//go:build (lang_ar || lang_all) && (country_africa || country_all || country_na || country_southern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNamibia.RegisterName(xlanguage.Arabic, "ناميبيا")
	dataNamibia.RegisterOfficialName(xlanguage.Arabic, "جمهورية ناميبيا")
	dataNamibia.RegisterCapital(xlanguage.Arabic, "ويندهوك")
}
