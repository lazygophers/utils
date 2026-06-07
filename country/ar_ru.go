//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArgentina.RegisterName(xlanguage.Russian, "Аргентина")
	dataArgentina.RegisterOfficialName(xlanguage.Russian, "Аргентинская Республика")
	dataArgentina.RegisterCapital(xlanguage.Russian, "Буэнос-Айрес")
}
