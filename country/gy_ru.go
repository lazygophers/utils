//go:build (lang_ru || lang_all) && (country_all || country_americas || country_gy || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuyana.RegisterName(xlanguage.Russian, "Гайана")
	dataGuyana.RegisterOfficialName(xlanguage.Russian, "Кооперативная Республика Гайана")
	dataGuyana.RegisterCapital(xlanguage.Russian, "Джорджтаун")
}
