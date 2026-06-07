//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_ke)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKenya.RegisterName(xlanguage.Arabic, "كينيا")
	dataKenya.RegisterOfficialName(xlanguage.Arabic, "جمهورية كينيا")
	dataKenya.RegisterCapital(xlanguage.Arabic, "نيروبي")
}
