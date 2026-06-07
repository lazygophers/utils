//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	RegionEasternAsia.RegisterName(xlanguage.Korean, "동아시아")
	RegionSouthEasternAsia.RegisterName(xlanguage.Korean, "동남아시아")
	RegionSouthernAsia.RegisterName(xlanguage.Korean, "남아시아")
	RegionWesternAsia.RegisterName(xlanguage.Korean, "서아시아")
	RegionCentralAsia.RegisterName(xlanguage.Korean, "중앙아시아")
	RegionEasternEurope.RegisterName(xlanguage.Korean, "동유럽")
	RegionNorthernEurope.RegisterName(xlanguage.Korean, "북유럽")
	RegionSouthernEurope.RegisterName(xlanguage.Korean, "남유럽")
	RegionWesternEurope.RegisterName(xlanguage.Korean, "서유럽")
	RegionNorthernAfrica.RegisterName(xlanguage.Korean, "북아프리카")
	RegionEasternAfrica.RegisterName(xlanguage.Korean, "동아프리카")
	RegionMiddleAfrica.RegisterName(xlanguage.Korean, "중앙아프리카")
	RegionSouthernAfrica.RegisterName(xlanguage.Korean, "남아프리카")
	RegionWesternAfrica.RegisterName(xlanguage.Korean, "서아프리카")
	RegionNorthernAmerica.RegisterName(xlanguage.Korean, "북아메리카")
	RegionCentralAmerica.RegisterName(xlanguage.Korean, "중앙아메리카")
	RegionSouthAmerica.RegisterName(xlanguage.Korean, "남아메리카")
	RegionCaribbean.RegisterName(xlanguage.Korean, "카리브 지역")
	RegionAustraliaAndNewZealand.RegisterName(xlanguage.Korean, "오스트레일리아와 뉴질랜드")
	RegionMelanesia.RegisterName(xlanguage.Korean, "멜라네시아")
	RegionMicronesia.RegisterName(xlanguage.Korean, "미크로네시아")
	RegionPolynesia.RegisterName(xlanguage.Korean, "폴리네시아")
	RegionAntarctic.RegisterName(xlanguage.Korean, "남극 지역")
}
