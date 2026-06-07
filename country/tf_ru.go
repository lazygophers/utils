//go:build (lang_ru || lang_all) && (country_all || country_antarctic || country_tf)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchSouthernTerritories.RegisterName(xlanguage.Russian, "Французские Южные и Антарктические территории")
	dataFrenchSouthernTerritories.RegisterOfficialName(xlanguage.Russian, "Территория Французских Южных и Антарктических земель")
	dataFrenchSouthernTerritories.RegisterCapital(xlanguage.Russian, "Порт-о-Франсе")
}
