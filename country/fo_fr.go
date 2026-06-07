//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFaroeIslands.RegisterName(xlanguage.French, "Îles Féroé")
	dataFaroeIslands.RegisterOfficialName(xlanguage.French, "Îles Féroé")
	dataFaroeIslands.RegisterCapital(xlanguage.French, "Tórshavn")
}
