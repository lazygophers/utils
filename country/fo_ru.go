//go:build (lang_ru || lang_all) && (country_all || country_europe || country_fo || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFaroeIslands.RegisterName(xlanguage.Russian, "Фарерские острова")
	dataFaroeIslands.RegisterOfficialName(xlanguage.Russian, "Фарерские острова")
	dataFaroeIslands.RegisterCapital(xlanguage.Russian, "Торсхавн")
}
