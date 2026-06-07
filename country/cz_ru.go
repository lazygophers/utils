//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCzechia.RegisterName(xlanguage.Russian, "Чехия")
	dataCzechia.RegisterOfficialName(xlanguage.Russian, "Чешская Республика")
	dataCzechia.RegisterCapital(xlanguage.Russian, "Прага")
}
