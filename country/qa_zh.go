//go:build country_all || country_asia || country_qa || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataQatar.RegisterName(xlanguage.Chinese, "卡塔尔")
	dataQatar.RegisterOfficialName(xlanguage.Chinese, "卡塔尔国")
	dataQatar.RegisterCapital(xlanguage.Chinese, "多哈")
}
