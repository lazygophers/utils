//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTanzania.RegisterName(xlanguage.Russian, "Танзания")
	dataTanzania.RegisterOfficialName(xlanguage.Russian, "Объединённая Республика Танзания")
	dataTanzania.RegisterCapital(xlanguage.Russian, "Додома")
}
