//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_tz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTanzania.RegisterName(xlanguage.Russian, "Танзания")
	dataTanzania.RegisterOfficialName(xlanguage.Russian, "Объединённая Республика Танзания")
	dataTanzania.RegisterCapital(xlanguage.Russian, "Додома")
}
