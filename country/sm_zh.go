//go:build country_all || country_europe || country_sm || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSanMarino.RegisterName(xlanguage.Chinese, "圣马力诺")
	dataSanMarino.RegisterOfficialName(xlanguage.Chinese, "圣马力诺共和国")
	dataSanMarino.RegisterCapital(xlanguage.Chinese, "圣马力诺")
}
