//go:build (lang_ar || lang_all) && (country_all || country_asia || country_central_asia || country_kg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKyrgyzstan.RegisterName(xlanguage.Arabic, "قيرغيزستان")
	dataKyrgyzstan.RegisterOfficialName(xlanguage.Arabic, "جمهورية قيرغيزستان")
	dataKyrgyzstan.RegisterCapital(xlanguage.Arabic, "بيشكك")
}
