//go:build country_all || country_europe || country_me || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontenegro.RegisterName(xlanguage.Chinese, "黑山")
	dataMontenegro.RegisterOfficialName(xlanguage.Chinese, "黑山")
	dataMontenegro.RegisterCapital(xlanguage.Chinese, "波德戈里察")
}
