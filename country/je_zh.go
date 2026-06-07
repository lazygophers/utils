//go:build country_all || country_europe || country_je || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJersey.RegisterName(xlanguage.Chinese, "泽西")
	dataJersey.RegisterOfficialName(xlanguage.Chinese, "泽西行政区")
	dataJersey.RegisterCapital(xlanguage.Chinese, "圣赫利尔")
}
