//go:build (lang_ru || lang_all) && (country_all || country_europe || country_lt || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLithuania.RegisterName(xlanguage.Russian, "Литва")
	dataLithuania.RegisterOfficialName(xlanguage.Russian, "Литовская Республика")
	dataLithuania.RegisterCapital(xlanguage.Russian, "Вильнюс")
}
