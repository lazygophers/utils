//go:build (lang_ar || lang_all) && (country_all || country_asia || country_id || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndonesia.RegisterName(xlanguage.Arabic, "إندونيسيا")
	dataIndonesia.RegisterOfficialName(xlanguage.Arabic, "جمهورية إندونيسيا")
	dataIndonesia.RegisterCapital(xlanguage.Arabic, "جاكرتا")
}
