//go:build country_all || country_europe || country_lt || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLithuania.RegisterName(xlanguage.Chinese, "立陶宛")
	dataLithuania.RegisterOfficialName(xlanguage.Chinese, "立陶宛共和国")
	dataLithuania.RegisterCapital(xlanguage.Chinese, "维尔纽斯")
}
