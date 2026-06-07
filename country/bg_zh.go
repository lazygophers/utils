//go:build country_all || country_bg || country_eastern_europe || country_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBulgaria.RegisterName(xlanguage.Chinese, "保加利亚")
	dataBulgaria.RegisterOfficialName(xlanguage.Chinese, "保加利亚共和国")
	dataBulgaria.RegisterCapital(xlanguage.Chinese, "索非亚")
}
