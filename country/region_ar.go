//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	RegionEasternAsia.RegisterName(xlanguage.Arabic, "شرق آسيا")
	RegionSouthEasternAsia.RegisterName(xlanguage.Arabic, "جنوب شرق آسيا")
	RegionSouthernAsia.RegisterName(xlanguage.Arabic, "جنوب آسيا")
	RegionWesternAsia.RegisterName(xlanguage.Arabic, "غرب آسيا")
	RegionCentralAsia.RegisterName(xlanguage.Arabic, "آسيا الوسطى")
	RegionEasternEurope.RegisterName(xlanguage.Arabic, "أوروبا الشرقية")
	RegionNorthernEurope.RegisterName(xlanguage.Arabic, "أوروبا الشمالية")
	RegionSouthernEurope.RegisterName(xlanguage.Arabic, "أوروبا الجنوبية")
	RegionWesternEurope.RegisterName(xlanguage.Arabic, "أوروبا الغربية")
	RegionNorthernAfrica.RegisterName(xlanguage.Arabic, "شمال أفريقيا")
	RegionEasternAfrica.RegisterName(xlanguage.Arabic, "شرق أفريقيا")
	RegionMiddleAfrica.RegisterName(xlanguage.Arabic, "وسط أفريقيا")
	RegionSouthernAfrica.RegisterName(xlanguage.Arabic, "الجنوب الأفريقي")
	RegionWesternAfrica.RegisterName(xlanguage.Arabic, "غرب أفريقيا")
	RegionNorthernAmerica.RegisterName(xlanguage.Arabic, "أمريكا الشمالية")
	RegionCentralAmerica.RegisterName(xlanguage.Arabic, "أمريكا الوسطى")
	RegionSouthAmerica.RegisterName(xlanguage.Arabic, "أمريكا الجنوبية")
	RegionCaribbean.RegisterName(xlanguage.Arabic, "البحر الكاريبي")
	RegionAustraliaAndNewZealand.RegisterName(xlanguage.Arabic, "أستراليا ونيوزيلندا")
	RegionMelanesia.RegisterName(xlanguage.Arabic, "ميلانيزيا")
	RegionMicronesia.RegisterName(xlanguage.Arabic, "ميكرونيزيا")
	RegionPolynesia.RegisterName(xlanguage.Arabic, "بولينيزيا")
	RegionAntarctic.RegisterName(xlanguage.Arabic, "المنطقة القطبية الجنوبية")
}
