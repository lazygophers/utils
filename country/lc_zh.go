//go:build country_all || country_americas || country_caribbean || country_lc

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintLucia.RegisterName(xlanguage.Chinese, "圣卢西亚")
	dataSaintLucia.RegisterOfficialName(xlanguage.Chinese, "圣卢西亚")
	dataSaintLucia.RegisterCapital(xlanguage.Chinese, "卡斯特里")
}
