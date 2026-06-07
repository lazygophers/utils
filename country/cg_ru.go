//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCongo.RegisterName(xlanguage.Russian, "Республика Конго")
	dataCongo.RegisterOfficialName(xlanguage.Russian, "Республика Конго")
	dataCongo.RegisterCapital(xlanguage.Russian, "Браззавиль")
}
