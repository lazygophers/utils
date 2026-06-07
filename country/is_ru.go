//go:build (lang_ru || lang_all) && (country_all || country_europe || country_is || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIceland.RegisterName(xlanguage.Russian, "Исландия")
	dataIceland.RegisterOfficialName(xlanguage.Russian, "Исландия")
	dataIceland.RegisterCapital(xlanguage.Russian, "Рейкьявик")
}
