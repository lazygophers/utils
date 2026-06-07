//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaboVerde.RegisterName(xlanguage.Russian, "Кабо-Верде")
	dataCaboVerde.RegisterOfficialName(xlanguage.Russian, "Республика Кабо-Верде")
	dataCaboVerde.RegisterCapital(xlanguage.Russian, "Прая")
}
