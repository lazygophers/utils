//go:build country_all || country_europe || country_lu || country_western_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLuxembourg.RegisterName(xlanguage.Chinese, "卢森堡")
	dataLuxembourg.RegisterOfficialName(xlanguage.Chinese, "卢森堡大公国")
	dataLuxembourg.RegisterCapital(xlanguage.Chinese, "卢森堡市")
}
