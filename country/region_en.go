package country

import xlanguage "golang.org/x/text/language"

// English is the canonical fallback; explicitly registered so that
// Region.NameIn(English) hits the map rather than the package-level default.
func init() {
	RegionEasternAsia.RegisterName(xlanguage.English, "Eastern Asia")
	RegionSouthEasternAsia.RegisterName(xlanguage.English, "South-eastern Asia")
	RegionSouthernAsia.RegisterName(xlanguage.English, "Southern Asia")
	RegionWesternAsia.RegisterName(xlanguage.English, "Western Asia")
	RegionCentralAsia.RegisterName(xlanguage.English, "Central Asia")
	RegionEasternEurope.RegisterName(xlanguage.English, "Eastern Europe")
	RegionNorthernEurope.RegisterName(xlanguage.English, "Northern Europe")
	RegionSouthernEurope.RegisterName(xlanguage.English, "Southern Europe")
	RegionWesternEurope.RegisterName(xlanguage.English, "Western Europe")
	RegionNorthernAfrica.RegisterName(xlanguage.English, "Northern Africa")
	RegionEasternAfrica.RegisterName(xlanguage.English, "Eastern Africa")
	RegionMiddleAfrica.RegisterName(xlanguage.English, "Middle Africa")
	RegionSouthernAfrica.RegisterName(xlanguage.English, "Southern Africa")
	RegionWesternAfrica.RegisterName(xlanguage.English, "Western Africa")
	RegionNorthernAmerica.RegisterName(xlanguage.English, "Northern America")
	RegionCentralAmerica.RegisterName(xlanguage.English, "Central America")
	RegionSouthAmerica.RegisterName(xlanguage.English, "South America")
	RegionCaribbean.RegisterName(xlanguage.English, "Caribbean")
	RegionAustraliaAndNewZealand.RegisterName(xlanguage.English, "Australia and New Zealand")
	RegionMelanesia.RegisterName(xlanguage.English, "Melanesia")
	RegionMicronesia.RegisterName(xlanguage.English, "Micronesia")
	RegionPolynesia.RegisterName(xlanguage.English, "Polynesia")
	RegionAntarctic.RegisterName(xlanguage.English, "Antarctic")
}
