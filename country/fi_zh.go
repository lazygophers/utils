//go:build country_all || country_europe || country_fi || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFinland.RegisterName(xlanguage.Chinese, "芬兰")
	dataFinland.RegisterOfficialName(xlanguage.Chinese, "芬兰共和国")
	dataFinland.RegisterCapital(xlanguage.Chinese, "赫尔辛基")
}
