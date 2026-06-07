//go:build (lang_fr || lang_all) && (country_all || country_asia || country_central_asia || country_kg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKyrgyzstan.RegisterName(xlanguage.French, "Kirghizistan")
	dataKyrgyzstan.RegisterOfficialName(xlanguage.French, "République kirghize")
	dataKyrgyzstan.RegisterCapital(xlanguage.French, "Bichkek")
}
