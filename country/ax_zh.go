//go:build country_all || country_ax || country_europe || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlandIslands.RegisterName(xlanguage.Chinese, "奥兰群岛")
	dataAlandIslands.RegisterOfficialName(xlanguage.Chinese, "奥兰群岛")
	dataAlandIslands.RegisterCapital(xlanguage.Chinese, "玛丽港")
}
