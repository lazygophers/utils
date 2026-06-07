//go:build country_africa || country_all || country_eastern_africa || country_ke

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKenya.RegisterName(xlanguage.Chinese, "肯尼亚")
	dataKenya.RegisterOfficialName(xlanguage.Chinese, "肯尼亚共和国")
	dataKenya.RegisterCapital(xlanguage.Chinese, "内罗毕")
}
