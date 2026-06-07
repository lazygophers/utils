//go:build (lang_ru || lang_all) && (country_all || country_eastern_europe || country_europe || country_hu)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHungary.RegisterName(xlanguage.Russian, "Венгрия")
	dataHungary.RegisterOfficialName(xlanguage.Russian, "Венгрия")
	dataHungary.RegisterCapital(xlanguage.Russian, "Будапешт")
}
