//go:build (lang_ar || lang_all) && (country_all || country_at || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustria.RegisterName(xlanguage.Arabic, "النمسا")
	dataAustria.RegisterOfficialName(xlanguage.Arabic, "جمهورية النمسا")
	dataAustria.RegisterCapital(xlanguage.Arabic, "فيينا")
}
