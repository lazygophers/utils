//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	RegionEasternAsia.RegisterName(xlanguage.Spanish, "Asia Oriental")
	RegionSouthEasternAsia.RegisterName(xlanguage.Spanish, "Sudeste Asiático")
	RegionSouthernAsia.RegisterName(xlanguage.Spanish, "Asia Meridional")
	RegionWesternAsia.RegisterName(xlanguage.Spanish, "Asia Occidental")
	RegionCentralAsia.RegisterName(xlanguage.Spanish, "Asia Central")
	RegionEasternEurope.RegisterName(xlanguage.Spanish, "Europa Oriental")
	RegionNorthernEurope.RegisterName(xlanguage.Spanish, "Europa Septentrional")
	RegionSouthernEurope.RegisterName(xlanguage.Spanish, "Europa Meridional")
	RegionWesternEurope.RegisterName(xlanguage.Spanish, "Europa Occidental")
	RegionNorthernAfrica.RegisterName(xlanguage.Spanish, "África Septentrional")
	RegionEasternAfrica.RegisterName(xlanguage.Spanish, "África Oriental")
	RegionMiddleAfrica.RegisterName(xlanguage.Spanish, "África Central")
	RegionSouthernAfrica.RegisterName(xlanguage.Spanish, "África Meridional")
	RegionWesternAfrica.RegisterName(xlanguage.Spanish, "África Occidental")
	RegionNorthernAmerica.RegisterName(xlanguage.Spanish, "América del Norte")
	RegionCentralAmerica.RegisterName(xlanguage.Spanish, "América Central")
	RegionSouthAmerica.RegisterName(xlanguage.Spanish, "América del Sur")
	RegionCaribbean.RegisterName(xlanguage.Spanish, "Caribe")
	RegionAustraliaAndNewZealand.RegisterName(xlanguage.Spanish, "Australia y Nueva Zelanda")
	RegionMelanesia.RegisterName(xlanguage.Spanish, "Melanesia")
	RegionMicronesia.RegisterName(xlanguage.Spanish, "Micronesia")
	RegionPolynesia.RegisterName(xlanguage.Spanish, "Polinesia")
	RegionAntarctic.RegisterName(xlanguage.Spanish, "Antártida")
}
