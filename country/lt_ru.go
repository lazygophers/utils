//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLithuania.RegisterName(xlanguage.Russian, "Литва")
	dataLithuania.RegisterOfficialName(xlanguage.Russian, "Литовская Республика")
	dataLithuania.RegisterCapital(xlanguage.Russian, "Вильнюс")
}
