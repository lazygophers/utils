package country

import xlanguage "golang.org/x/text/language"

func init() {
	RegionEasternAsia.RegisterName(xlanguage.Chinese, "东亚")
	RegionSouthEasternAsia.RegisterName(xlanguage.Chinese, "东南亚")
	RegionSouthernAsia.RegisterName(xlanguage.Chinese, "南亚")
	RegionWesternAsia.RegisterName(xlanguage.Chinese, "西亚")
	RegionCentralAsia.RegisterName(xlanguage.Chinese, "中亚")
	RegionEasternEurope.RegisterName(xlanguage.Chinese, "东欧")
	RegionNorthernEurope.RegisterName(xlanguage.Chinese, "北欧")
	RegionSouthernEurope.RegisterName(xlanguage.Chinese, "南欧")
	RegionWesternEurope.RegisterName(xlanguage.Chinese, "西欧")
	RegionNorthernAfrica.RegisterName(xlanguage.Chinese, "北非")
	RegionEasternAfrica.RegisterName(xlanguage.Chinese, "东非")
	RegionMiddleAfrica.RegisterName(xlanguage.Chinese, "中非")
	RegionSouthernAfrica.RegisterName(xlanguage.Chinese, "南部非洲")
	RegionWesternAfrica.RegisterName(xlanguage.Chinese, "西非")
	RegionNorthernAmerica.RegisterName(xlanguage.Chinese, "北美洲")
	RegionCentralAmerica.RegisterName(xlanguage.Chinese, "中美洲")
	RegionSouthAmerica.RegisterName(xlanguage.Chinese, "南美洲")
	RegionCaribbean.RegisterName(xlanguage.Chinese, "加勒比地区")
	RegionAustraliaAndNewZealand.RegisterName(xlanguage.Chinese, "澳大利亚和新西兰")
	RegionMelanesia.RegisterName(xlanguage.Chinese, "美拉尼西亚")
	RegionMicronesia.RegisterName(xlanguage.Chinese, "密克罗尼西亚")
	RegionPolynesia.RegisterName(xlanguage.Chinese, "波利尼西亚")
	RegionAntarctic.RegisterName(xlanguage.Chinese, "南极地区")
}
