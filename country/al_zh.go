//go:build country_al || country_all || country_europe || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlbania.RegisterName(xlanguage.Chinese, "阿尔巴尼亚")
	dataAlbania.RegisterOfficialName(xlanguage.Chinese, "阿尔巴尼亚共和国")
	dataAlbania.RegisterCapital(xlanguage.Chinese, "地拉那")
}
