//go:build country_africa || country_all || country_ga || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGabon.RegisterName(xlanguage.Chinese, "加蓬")
	dataGabon.RegisterOfficialName(xlanguage.Chinese, "加蓬共和国")
	dataGabon.RegisterCapital(xlanguage.Chinese, "利伯维尔")
}
