//go:build (lang_ru || lang_all) && (country_all || country_americas || country_aw || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAruba.RegisterName(xlanguage.Russian, "Аруба")
	dataAruba.RegisterOfficialName(xlanguage.Russian, "Аруба")
	dataAruba.RegisterCapital(xlanguage.Russian, "Ораньестад")
}
