//go:build country_all || country_europe || country_is || country_northern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIceland.RegisterName(xlanguage.Chinese, "冰岛")
	dataIceland.RegisterOfficialName(xlanguage.Chinese, "冰岛共和国")
	dataIceland.RegisterCapital(xlanguage.Chinese, "雷克雅未克")
}
