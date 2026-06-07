//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEgypt.RegisterName(xlanguage.Russian, "Египет")
	dataEgypt.RegisterOfficialName(xlanguage.Russian, "Арабская Республика Египет")
	dataEgypt.RegisterCapital(xlanguage.Russian, "Каир")
}
