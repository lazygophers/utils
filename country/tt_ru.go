//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTrinidadAndTobago.RegisterName(xlanguage.Russian, "Тринидад и Тобаго")
	dataTrinidadAndTobago.RegisterOfficialName(xlanguage.Russian, "Республика Тринидад и Тобаго")
	dataTrinidadAndTobago.RegisterCapital(xlanguage.Russian, "Порт-оф-Спейн")
}
