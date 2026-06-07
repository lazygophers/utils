//go:build country_all || country_americas || country_caribbean || country_vi

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsVirginIslands.RegisterName(xlanguage.Chinese, "美属维尔京群岛")
	dataUsVirginIslands.RegisterOfficialName(xlanguage.Chinese, "美属维尔京群岛")
	dataUsVirginIslands.RegisterCapital(xlanguage.Chinese, "夏洛特阿马利亚")
}
