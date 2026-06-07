//go:build (lang_ar || lang_all) && (country_africa || country_all || country_ao || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAngola.RegisterName(xlanguage.Arabic, "أنغولا")
	dataAngola.RegisterOfficialName(xlanguage.Arabic, "جمهورية أنغولا")
	dataAngola.RegisterCapital(xlanguage.Arabic, "لواندا")
}
