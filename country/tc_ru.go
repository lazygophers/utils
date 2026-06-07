//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurksAndCaicosIslands.RegisterName(xlanguage.Russian, "Тёркс и Кайкос")
	dataTurksAndCaicosIslands.RegisterOfficialName(xlanguage.Russian, "Острова Тёркс и Кайкос")
	dataTurksAndCaicosIslands.RegisterCapital(xlanguage.Russian, "Коберн-Таун")
}
