//go:build (lang_ru || lang_all) && (country_all || country_asia || country_south_eastern_asia || country_tl)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTimorLeste.RegisterName(xlanguage.Russian, "Восточный Тимор")
	dataTimorLeste.RegisterOfficialName(xlanguage.Russian, "Демократическая Республика Тимор-Лешти")
	dataTimorLeste.RegisterCapital(xlanguage.Russian, "Дили")
}
