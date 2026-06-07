//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMyanmar.RegisterName(xlanguage.Russian, "Мьянма")
	dataMyanmar.RegisterOfficialName(xlanguage.Russian, "Республика Союз Мьянма")
	dataMyanmar.RegisterCapital(xlanguage.Russian, "Нейпьидо")
}
