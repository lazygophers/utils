//go:build (lang_ar || lang_all) && (country_africa || country_all || country_gn || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuinea.RegisterName(xlanguage.Arabic, "غينيا")
	dataGuinea.RegisterOfficialName(xlanguage.Arabic, "جمهورية غينيا")
	dataGuinea.RegisterCapital(xlanguage.Arabic, "كوناكري")
}
