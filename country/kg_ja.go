//go:build (lang_ja || lang_all) && (country_all || country_asia || country_central_asia || country_kg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKyrgyzstan.RegisterName(xlanguage.Japanese, "キルギス")
	dataKyrgyzstan.RegisterOfficialName(xlanguage.Japanese, "キルギス共和国")
	dataKyrgyzstan.RegisterCapital(xlanguage.Japanese, "ビシュケク")
}
