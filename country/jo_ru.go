//go:build (lang_ru || lang_all) && (country_all || country_asia || country_jo || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJordan.RegisterName(xlanguage.Russian, "Иордания")
	dataJordan.RegisterOfficialName(xlanguage.Russian, "Иорданское Хашимитское Королевство")
	dataJordan.RegisterCapital(xlanguage.Russian, "Амман")
}
