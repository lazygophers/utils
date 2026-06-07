//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eg || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEgypt.RegisterName(xlanguage.Russian, "Египет")
	dataEgypt.RegisterOfficialName(xlanguage.Russian, "Арабская Республика Египет")
	dataEgypt.RegisterCapital(xlanguage.Russian, "Каир")
}
