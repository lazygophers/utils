//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIran.RegisterName(xlanguage.Russian, "Иран")
	dataIran.RegisterOfficialName(xlanguage.Russian, "Исламская Республика Иран")
	dataIran.RegisterCapital(xlanguage.Russian, "Тегеран")
}
