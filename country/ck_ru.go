//go:build (lang_ru || lang_all) && (country_all || country_ck || country_oceania || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCookIslands.RegisterName(xlanguage.Russian, "Острова Кука")
	dataCookIslands.RegisterOfficialName(xlanguage.Russian, "Острова Кука")
	dataCookIslands.RegisterCapital(xlanguage.Russian, "Аваруа")
}
