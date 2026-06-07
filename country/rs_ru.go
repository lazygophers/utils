//go:build (lang_ru || lang_all) && (country_all || country_europe || country_rs || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSerbia.RegisterName(xlanguage.Russian, "Сербия")
	dataSerbia.RegisterOfficialName(xlanguage.Russian, "Республика Сербия")
	dataSerbia.RegisterCapital(xlanguage.Russian, "Белград")
}
