//go:build country_all || country_europe || country_pt || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPortugal.RegisterName(xlanguage.Chinese, "葡萄牙")
	dataPortugal.RegisterOfficialName(xlanguage.Chinese, "葡萄牙共和国")
	dataPortugal.RegisterCapital(xlanguage.Chinese, "里斯本")
}
