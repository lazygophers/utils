//go:build country_all || country_asia || country_central_asia || country_kg

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKyrgyzstan.RegisterName(xlanguage.Russian, "Киргизия")
	dataKyrgyzstan.RegisterOfficialName(xlanguage.Russian, "Киргизская Республика")
	dataKyrgyzstan.RegisterCapital(xlanguage.Russian, "Бишкек")
}
