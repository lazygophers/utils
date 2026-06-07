//go:build (lang_ru || lang_all) && (country_all || country_europe || country_gr || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreece.RegisterName(xlanguage.Russian, "Греция")
	dataGreece.RegisterOfficialName(xlanguage.Russian, "Греческая Республика")
	dataGreece.RegisterCapital(xlanguage.Russian, "Афины")
}
