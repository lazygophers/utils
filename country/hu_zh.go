//go:build country_all || country_eastern_europe || country_europe || country_hu

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHungary.RegisterName(xlanguage.Chinese, "匈牙利")
	dataHungary.RegisterOfficialName(xlanguage.Chinese, "匈牙利")
	dataHungary.RegisterCapital(xlanguage.Chinese, "布达佩斯")
}
