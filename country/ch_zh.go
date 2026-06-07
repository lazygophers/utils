//go:build country_all || country_ch || country_europe || country_western_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSwitzerland.RegisterName(xlanguage.Chinese, "瑞士")
	dataSwitzerland.RegisterOfficialName(xlanguage.Chinese, "瑞士联邦")
	dataSwitzerland.RegisterCapital(xlanguage.Chinese, "伯尔尼")
}
