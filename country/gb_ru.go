//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedKingdom.RegisterName(xlanguage.Russian, "Великобритания")
	dataUnitedKingdom.RegisterOfficialName(xlanguage.Russian, "Соединённое Королевство Великобритании и Северной Ирландии")
	dataUnitedKingdom.RegisterCapital(xlanguage.Russian, "Лондон")
}
