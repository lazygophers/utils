//go:build country_africa || country_all || country_na || country_southern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNamibia.RegisterName(xlanguage.Chinese, "纳米比亚")
	dataNamibia.RegisterOfficialName(xlanguage.Chinese, "纳米比亚共和国")
	dataNamibia.RegisterCapital(xlanguage.Chinese, "温得和克")
}
