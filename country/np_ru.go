//go:build (lang_ru || lang_all) && (country_all || country_asia || country_np || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNepal.RegisterName(xlanguage.Russian, "Непал")
	dataNepal.RegisterOfficialName(xlanguage.Russian, "Федеративная Демократическая Республика Непал")
	dataNepal.RegisterCapital(xlanguage.Russian, "Катманду")
}
