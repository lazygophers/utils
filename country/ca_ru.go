//go:build (lang_ru || lang_all) && (country_all || country_americas || country_ca || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCanada.RegisterName(xlanguage.Russian, "Канада")
	dataCanada.RegisterOfficialName(xlanguage.Russian, "Канада")
	dataCanada.RegisterCapital(xlanguage.Russian, "Оттава")
}
