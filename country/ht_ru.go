//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHaiti.RegisterName(xlanguage.Russian, "Гаити")
	dataHaiti.RegisterOfficialName(xlanguage.Russian, "Республика Гаити")
	dataHaiti.RegisterCapital(xlanguage.Russian, "Порт-о-Пренс")
}
