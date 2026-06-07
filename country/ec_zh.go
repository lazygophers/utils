//go:build country_all || country_americas || country_ec || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEcuador.RegisterName(xlanguage.Chinese, "厄瓜多尔")
	dataEcuador.RegisterOfficialName(xlanguage.Chinese, "厄瓜多尔共和国")
	dataEcuador.RegisterCapital(xlanguage.Chinese, "基多")
}
