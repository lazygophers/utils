//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFaroeIslands.RegisterName(xlanguage.Russian, "Фарерские острова")
	dataFaroeIslands.RegisterOfficialName(xlanguage.Russian, "Фарерские острова")
	dataFaroeIslands.RegisterCapital(xlanguage.Russian, "Торсхавн")
}
