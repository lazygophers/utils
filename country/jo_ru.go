//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJordan.RegisterName(xlanguage.Russian, "Иордания")
	dataJordan.RegisterOfficialName(xlanguage.Russian, "Иорданское Хашимитское Королевство")
	dataJordan.RegisterCapital(xlanguage.Russian, "Амман")
}
