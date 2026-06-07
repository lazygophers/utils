//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCocosKeelingIslands.RegisterName(xlanguage.Russian, "Кокосовые острова")
	dataCocosKeelingIslands.RegisterOfficialName(xlanguage.Russian, "Территория Кокосовые (Килинг) острова")
	dataCocosKeelingIslands.RegisterCapital(xlanguage.Russian, "Уэст-Айленд")
}
