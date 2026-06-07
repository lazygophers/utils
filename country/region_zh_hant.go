//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

var regionZhHant = xlanguage.MustParse("zh-Hant")

func init() {
	RegionEasternAsia.RegisterName(regionZhHant, "東亞")
	RegionSouthEasternAsia.RegisterName(regionZhHant, "東南亞")
	RegionSouthernAsia.RegisterName(regionZhHant, "南亞")
	RegionWesternAsia.RegisterName(regionZhHant, "西亞")
	RegionCentralAsia.RegisterName(regionZhHant, "中亞")
	RegionEasternEurope.RegisterName(regionZhHant, "東歐")
	RegionNorthernEurope.RegisterName(regionZhHant, "北歐")
	RegionSouthernEurope.RegisterName(regionZhHant, "南歐")
	RegionWesternEurope.RegisterName(regionZhHant, "西歐")
	RegionNorthernAfrica.RegisterName(regionZhHant, "北非")
	RegionEasternAfrica.RegisterName(regionZhHant, "東非")
	RegionMiddleAfrica.RegisterName(regionZhHant, "中非")
	RegionSouthernAfrica.RegisterName(regionZhHant, "南部非洲")
	RegionWesternAfrica.RegisterName(regionZhHant, "西非")
	RegionNorthernAmerica.RegisterName(regionZhHant, "北美洲")
	RegionCentralAmerica.RegisterName(regionZhHant, "中美洲")
	RegionSouthAmerica.RegisterName(regionZhHant, "南美洲")
	RegionCaribbean.RegisterName(regionZhHant, "加勒比地區")
	RegionAustraliaAndNewZealand.RegisterName(regionZhHant, "澳大利亞和紐西蘭")
	RegionMelanesia.RegisterName(regionZhHant, "美拉尼西亞")
	RegionMicronesia.RegisterName(regionZhHant, "密克羅尼西亞")
	RegionPolynesia.RegisterName(regionZhHant, "玻里尼西亞")
	RegionAntarctic.RegisterName(regionZhHant, "南極地區")
}
