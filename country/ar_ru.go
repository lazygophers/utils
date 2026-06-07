//go:build (lang_ru || lang_all) && (country_all || country_americas || country_ar || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArgentina.RegisterName(xlanguage.Russian, "Аргентина")
	dataArgentina.RegisterOfficialName(xlanguage.Russian, "Аргентинская Республика")
	dataArgentina.RegisterCapital(xlanguage.Russian, "Буэнос-Айрес")
}
