//go:build (lang_ru || lang_all) && (country_all || country_australia_and_new_zealand || country_nz || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewZealand.RegisterName(xlanguage.Russian, "Новая Зеландия")
	dataNewZealand.RegisterOfficialName(xlanguage.Russian, "Новая Зеландия")
	dataNewZealand.RegisterCapital(xlanguage.Russian, "Веллингтон")
}
