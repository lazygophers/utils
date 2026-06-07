//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChad.RegisterName(xlanguage.Russian, "Чад")
	dataChad.RegisterOfficialName(xlanguage.Russian, "Республика Чад")
	dataChad.RegisterCapital(xlanguage.Russian, "Нджамена")
}
