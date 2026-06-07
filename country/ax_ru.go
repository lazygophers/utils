//go:build (lang_ru || lang_all) && (country_all || country_ax || country_europe || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlandIslands.RegisterName(xlanguage.Russian, "Аландские острова")
	dataAlandIslands.RegisterOfficialName(xlanguage.Russian, "Аландские острова")
	dataAlandIslands.RegisterCapital(xlanguage.Russian, "Мариехамн")
}
