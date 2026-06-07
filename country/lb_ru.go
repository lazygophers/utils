//go:build (lang_ru || lang_all) && (country_all || country_asia || country_lb || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLebanon.RegisterName(xlanguage.Russian, "Ливан")
	dataLebanon.RegisterOfficialName(xlanguage.Russian, "Ливанская Республика")
	dataLebanon.RegisterCapital(xlanguage.Russian, "Бейрут")
}
