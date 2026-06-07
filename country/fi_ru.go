//go:build (lang_ru || lang_all) && (country_all || country_europe || country_fi || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFinland.RegisterName(xlanguage.Russian, "Финляндия")
	dataFinland.RegisterOfficialName(xlanguage.Russian, "Финляндская Республика")
	dataFinland.RegisterCapital(xlanguage.Russian, "Хельсинки")
}
