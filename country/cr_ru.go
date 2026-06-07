//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCostaRica.RegisterName(xlanguage.Russian, "Коста-Рика")
	dataCostaRica.RegisterOfficialName(xlanguage.Russian, "Республика Коста-Рика")
	dataCostaRica.RegisterCapital(xlanguage.Russian, "Сан-Хосе")
}
