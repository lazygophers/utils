//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_tc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurksAndCaicosIslands.RegisterName(xlanguage.Russian, "Тёркс и Кайкос")
	dataTurksAndCaicosIslands.RegisterOfficialName(xlanguage.Russian, "Острова Тёркс и Кайкос")
	dataTurksAndCaicosIslands.RegisterCapital(xlanguage.Russian, "Коберн-Таун")
}
