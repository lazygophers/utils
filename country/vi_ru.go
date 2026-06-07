//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsVirginIslands.RegisterName(xlanguage.Russian, "Виргинские Острова США")
	dataUsVirginIslands.RegisterOfficialName(xlanguage.Russian, "Виргинские Острова Соединённых Штатов")
	dataUsVirginIslands.RegisterCapital(xlanguage.Russian, "Шарлотта-Амалия")
}
