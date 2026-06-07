//go:build (lang_ru || lang_all) && (country_all || country_europe || country_northern_europe || country_se)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSweden.RegisterName(xlanguage.Russian, "Швеция")
	dataSweden.RegisterOfficialName(xlanguage.Russian, "Королевство Швеция")
	dataSweden.RegisterCapital(xlanguage.Russian, "Стокгольм")
}
