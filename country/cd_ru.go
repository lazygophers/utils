//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDrCongo.RegisterName(xlanguage.Russian, "ДР Конго")
	dataDrCongo.RegisterOfficialName(xlanguage.Russian, "Демократическая Республика Конго")
	dataDrCongo.RegisterCapital(xlanguage.Russian, "Киншаса")
}
