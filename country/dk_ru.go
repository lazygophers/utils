//go:build (lang_ru || lang_all) && (country_all || country_dk || country_europe || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDenmark.RegisterName(xlanguage.Russian, "Дания")
	dataDenmark.RegisterOfficialName(xlanguage.Russian, "Королевство Дания")
	dataDenmark.RegisterCapital(xlanguage.Russian, "Копенгаген")
}
