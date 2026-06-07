//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTuvalu.RegisterName(xlanguage.Russian, "Тувалу")
	dataTuvalu.RegisterOfficialName(xlanguage.Russian, "Тувалу")
	dataTuvalu.RegisterCapital(xlanguage.Russian, "Фунафути")
}
