//go:build country_all || country_antarctic || country_tf

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchSouthernTerritories.RegisterName(xlanguage.French, "Terres australes et antarctiques françaises")
	dataFrenchSouthernTerritories.RegisterOfficialName(xlanguage.French, "Terres australes et antarctiques françaises")
	dataFrenchSouthernTerritories.RegisterCapital(xlanguage.French, "Saint-Pierre")
}
