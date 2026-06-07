//go:build country_all || country_europe || country_lv || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLatvia.RegisterName(xlanguage.Chinese, "拉脱维亚")
	dataLatvia.RegisterOfficialName(xlanguage.Chinese, "拉脱维亚共和国")
	dataLatvia.RegisterCapital(xlanguage.Chinese, "里加")
}
