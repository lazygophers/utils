//go:build country_all || country_asia || country_om || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataOman.RegisterName(xlanguage.Chinese, "阿曼")
	dataOman.RegisterOfficialName(xlanguage.Chinese, "阿曼苏丹国")
	dataOman.RegisterCapital(xlanguage.Chinese, "马斯喀特")
}
