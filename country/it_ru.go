//go:build (lang_ru || lang_all) && (country_all || country_europe || country_it || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataItaly.RegisterName(xlanguage.Russian, "Италия")
	dataItaly.RegisterOfficialName(xlanguage.Russian, "Итальянская Республика")
	dataItaly.RegisterCapital(xlanguage.Russian, "Рим")
}
