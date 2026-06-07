//go:build (lang_ru || lang_all) && (country_all || country_asia || country_id || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndonesia.RegisterName(xlanguage.Russian, "Индонезия")
	dataIndonesia.RegisterOfficialName(xlanguage.Russian, "Республика Индонезия")
	dataIndonesia.RegisterCapital(xlanguage.Russian, "Джакарта")
}
