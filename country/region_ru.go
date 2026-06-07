//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	RegionEasternAsia.RegisterName(xlanguage.Russian, "Восточная Азия")
	RegionSouthEasternAsia.RegisterName(xlanguage.Russian, "Юго-Восточная Азия")
	RegionSouthernAsia.RegisterName(xlanguage.Russian, "Южная Азия")
	RegionWesternAsia.RegisterName(xlanguage.Russian, "Западная Азия")
	RegionCentralAsia.RegisterName(xlanguage.Russian, "Центральная Азия")
	RegionEasternEurope.RegisterName(xlanguage.Russian, "Восточная Европа")
	RegionNorthernEurope.RegisterName(xlanguage.Russian, "Северная Европа")
	RegionSouthernEurope.RegisterName(xlanguage.Russian, "Южная Европа")
	RegionWesternEurope.RegisterName(xlanguage.Russian, "Западная Европа")
	RegionNorthernAfrica.RegisterName(xlanguage.Russian, "Северная Африка")
	RegionEasternAfrica.RegisterName(xlanguage.Russian, "Восточная Африка")
	RegionMiddleAfrica.RegisterName(xlanguage.Russian, "Центральная Африка")
	RegionSouthernAfrica.RegisterName(xlanguage.Russian, "Южная Африка")
	RegionWesternAfrica.RegisterName(xlanguage.Russian, "Западная Африка")
	RegionNorthernAmerica.RegisterName(xlanguage.Russian, "Северная Америка")
	RegionCentralAmerica.RegisterName(xlanguage.Russian, "Центральная Америка")
	RegionSouthAmerica.RegisterName(xlanguage.Russian, "Южная Америка")
	RegionCaribbean.RegisterName(xlanguage.Russian, "Карибский бассейн")
	RegionAustraliaAndNewZealand.RegisterName(xlanguage.Russian, "Австралия и Новая Зеландия")
	RegionMelanesia.RegisterName(xlanguage.Russian, "Меланезия")
	RegionMicronesia.RegisterName(xlanguage.Russian, "Микронезия")
	RegionPolynesia.RegisterName(xlanguage.Russian, "Полинезия")
	RegionAntarctic.RegisterName(xlanguage.Russian, "Антарктика")
}
