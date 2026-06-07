//go:build (lang_ru || lang_all) && (country_all || country_europe || country_no || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorway.RegisterName(xlanguage.Russian, "Норвегия")
	dataNorway.RegisterOfficialName(xlanguage.Russian, "Королевство Норвегия")
	dataNorway.RegisterCapital(xlanguage.Russian, "Осло")
}
