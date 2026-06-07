//go:build country_all || country_europe || country_fo || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFaroeIslands.RegisterName(xlanguage.Chinese, "法罗群岛")
	dataFaroeIslands.RegisterOfficialName(xlanguage.Chinese, "法罗群岛")
	dataFaroeIslands.RegisterCapital(xlanguage.Chinese, "托尔斯港")
}
