//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuadeloupe.RegisterName(xlanguage.Russian, "Гваделупа")
	dataGuadeloupe.RegisterOfficialName(xlanguage.Russian, "Гваделупа")
	dataGuadeloupe.RegisterCapital(xlanguage.Russian, "Бас-Тер")
}
