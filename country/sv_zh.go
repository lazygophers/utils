//go:build country_all || country_americas || country_central_america || country_sv

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataElSalvador.RegisterName(xlanguage.Chinese, "萨尔瓦多")
	dataElSalvador.RegisterOfficialName(xlanguage.Chinese, "萨尔瓦多共和国")
	dataElSalvador.RegisterCapital(xlanguage.Chinese, "圣萨尔瓦多")
}
