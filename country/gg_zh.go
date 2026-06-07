//go:build country_all || country_europe || country_gg || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuernsey.RegisterName(xlanguage.Chinese, "根西")
	dataGuernsey.RegisterOfficialName(xlanguage.Chinese, "根西行政区")
	dataGuernsey.RegisterCapital(xlanguage.Chinese, "圣彼得港")
}
