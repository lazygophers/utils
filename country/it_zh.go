//go:build country_all || country_europe || country_it || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataItaly.RegisterName(xlanguage.Chinese, "意大利")
	dataItaly.RegisterOfficialName(xlanguage.Chinese, "意大利共和国")
	dataItaly.RegisterCapital(xlanguage.Chinese, "罗马")
}
