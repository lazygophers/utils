//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthKorea.RegisterName(xlanguage.Russian, "КНДР")
	dataNorthKorea.RegisterOfficialName(xlanguage.Russian, "Корейская Народно-Демократическая Республика")
	dataNorthKorea.RegisterCapital(xlanguage.Russian, "Пхеньян")
}
