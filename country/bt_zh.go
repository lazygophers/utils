//go:build country_all || country_asia || country_bt || country_southern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBhutan.RegisterName(xlanguage.Chinese, "不丹")
	dataBhutan.RegisterOfficialName(xlanguage.Chinese, "不丹王国")
	dataBhutan.RegisterCapital(xlanguage.Chinese, "廷布")
}
