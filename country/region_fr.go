//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	RegionEasternAsia.RegisterName(xlanguage.French, "Asie de l'Est")
	RegionSouthEasternAsia.RegisterName(xlanguage.French, "Asie du Sud-Est")
	RegionSouthernAsia.RegisterName(xlanguage.French, "Asie du Sud")
	RegionWesternAsia.RegisterName(xlanguage.French, "Asie de l'Ouest")
	RegionCentralAsia.RegisterName(xlanguage.French, "Asie centrale")
	RegionEasternEurope.RegisterName(xlanguage.French, "Europe de l'Est")
	RegionNorthernEurope.RegisterName(xlanguage.French, "Europe du Nord")
	RegionSouthernEurope.RegisterName(xlanguage.French, "Europe du Sud")
	RegionWesternEurope.RegisterName(xlanguage.French, "Europe de l'Ouest")
	RegionNorthernAfrica.RegisterName(xlanguage.French, "Afrique du Nord")
	RegionEasternAfrica.RegisterName(xlanguage.French, "Afrique de l'Est")
	RegionMiddleAfrica.RegisterName(xlanguage.French, "Afrique centrale")
	RegionSouthernAfrica.RegisterName(xlanguage.French, "Afrique australe")
	RegionWesternAfrica.RegisterName(xlanguage.French, "Afrique de l'Ouest")
	RegionNorthernAmerica.RegisterName(xlanguage.French, "Amérique du Nord")
	RegionCentralAmerica.RegisterName(xlanguage.French, "Amérique centrale")
	RegionSouthAmerica.RegisterName(xlanguage.French, "Amérique du Sud")
	RegionCaribbean.RegisterName(xlanguage.French, "Caraïbes")
	RegionAustraliaAndNewZealand.RegisterName(xlanguage.French, "Australie et Nouvelle-Zélande")
	RegionMelanesia.RegisterName(xlanguage.French, "Mélanésie")
	RegionMicronesia.RegisterName(xlanguage.French, "Micronésie")
	RegionPolynesia.RegisterName(xlanguage.French, "Polynésie")
	RegionAntarctic.RegisterName(xlanguage.French, "Antarctique")
}
