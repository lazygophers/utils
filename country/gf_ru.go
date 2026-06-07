//go:build (lang_ru || lang_all) && (country_all || country_americas || country_gf || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchGuiana.RegisterName(xlanguage.Russian, "Французская Гвиана")
	dataFrenchGuiana.RegisterOfficialName(xlanguage.Russian, "Гвиана")
	dataFrenchGuiana.RegisterCapital(xlanguage.Russian, "Кайенна")
}
