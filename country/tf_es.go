//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchSouthernTerritories.RegisterName(xlanguage.Spanish, "Tierras Australes y Antárticas Francesas")
	dataFrenchSouthernTerritories.RegisterOfficialName(xlanguage.Spanish, "Tierras Australes y Antárticas Francesas")
	dataFrenchSouthernTerritories.RegisterCapital(xlanguage.Spanish, "Saint-Pierre")
}
