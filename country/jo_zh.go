//go:build country_all || country_asia || country_jo || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJordan.RegisterName(xlanguage.Chinese, "约旦")
	dataJordan.RegisterOfficialName(xlanguage.Chinese, "约旦哈希姆王国")
	dataJordan.RegisterCapital(xlanguage.Chinese, "安曼")
}
