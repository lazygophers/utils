//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEcuador.RegisterName(xlanguage.Russian, "Эквадор")
	dataEcuador.RegisterOfficialName(xlanguage.Russian, "Республика Эквадор")
	dataEcuador.RegisterCapital(xlanguage.Russian, "Кито")
}
