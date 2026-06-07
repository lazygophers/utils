//go:build country_all || country_europe || country_ie || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIreland.RegisterName(xlanguage.Chinese, "爱尔兰")
	dataIreland.RegisterOfficialName(xlanguage.Chinese, "爱尔兰")
	dataIreland.RegisterCapital(xlanguage.Chinese, "都柏林")
}
