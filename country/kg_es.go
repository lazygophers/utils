//go:build (lang_es || lang_all) && (country_all || country_asia || country_central_asia || country_kg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKyrgyzstan.RegisterName(xlanguage.Spanish, "Kirguistán")
	dataKyrgyzstan.RegisterOfficialName(xlanguage.Spanish, "República Kirguisa")
	dataKyrgyzstan.RegisterCapital(xlanguage.Spanish, "Biskek")
}
