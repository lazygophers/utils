//go:build country_all || country_ee || country_europe || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEstonia.RegisterName(xlanguage.Chinese, "爱沙尼亚")
	dataEstonia.RegisterOfficialName(xlanguage.Chinese, "爱沙尼亚共和国")
	dataEstonia.RegisterCapital(xlanguage.Chinese, "塔林")
}
