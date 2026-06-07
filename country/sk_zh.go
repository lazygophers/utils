//go:build country_all || country_eastern_europe || country_europe || country_sk

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovakia.RegisterName(xlanguage.Chinese, "斯洛伐克")
	dataSlovakia.RegisterOfficialName(xlanguage.Chinese, "斯洛伐克共和国")
	dataSlovakia.RegisterCapital(xlanguage.Chinese, "布拉迪斯拉发")
}
