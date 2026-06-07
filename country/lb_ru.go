//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLebanon.RegisterName(xlanguage.Russian, "Ливан")
	dataLebanon.RegisterOfficialName(xlanguage.Russian, "Ливанская Республика")
	dataLebanon.RegisterCapital(xlanguage.Russian, "Бейрут")
}
