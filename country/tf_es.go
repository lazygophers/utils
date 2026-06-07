//go:build (lang_es || lang_all) && (country_all || country_antarctic || country_tf)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchSouthernTerritories.RegisterName(xlanguage.Spanish, "Tierras Australes y Antárticas Francesas")
	dataFrenchSouthernTerritories.RegisterOfficialName(xlanguage.Spanish, "Tierras Australes y Antárticas Francesas")
	dataFrenchSouthernTerritories.RegisterCapital(xlanguage.Spanish, "Saint-Pierre")
}
