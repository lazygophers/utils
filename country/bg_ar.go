//go:build (lang_ar || lang_all) && (country_all || country_bg || country_eastern_europe || country_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBulgaria.RegisterName(xlanguage.Arabic, "بلغاريا")
	dataBulgaria.RegisterOfficialName(xlanguage.Arabic, "جمهورية بلغاريا")
	dataBulgaria.RegisterCapital(xlanguage.Arabic, "صوفيا")
}
