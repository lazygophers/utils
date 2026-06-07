//go:build (lang_ru || lang_all) && (country_all || country_ki || country_micronesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKiribati.RegisterName(xlanguage.Russian, "Кирибати")
	dataKiribati.RegisterOfficialName(xlanguage.Russian, "Республика Кирибати")
	dataKiribati.RegisterCapital(xlanguage.Russian, "Тарава")
}
