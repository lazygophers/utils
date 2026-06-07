//go:build (lang_ar || lang_all) && (country_all || country_asia || country_central_asia || country_kz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKazakhstan.RegisterName(xlanguage.Arabic, "كازاخستان")
	dataKazakhstan.RegisterOfficialName(xlanguage.Arabic, "جمهورية كازاخستان")
	dataKazakhstan.RegisterCapital(xlanguage.Arabic, "أستانا")
}
