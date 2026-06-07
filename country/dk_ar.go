//go:build (lang_ar || lang_all) && (country_all || country_dk || country_europe || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDenmark.RegisterName(xlanguage.Arabic, "الدنمارك")
	dataDenmark.RegisterOfficialName(xlanguage.Arabic, "مملكة الدنمارك")
	dataDenmark.RegisterCapital(xlanguage.Arabic, "كوبنهاغن")
}
