//go:build country_all || country_americas || country_aw || country_caribbean

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAruba.RegisterName(xlanguage.Chinese, "阿鲁巴")
	dataAruba.RegisterOfficialName(xlanguage.Chinese, "阿鲁巴")
	dataAruba.RegisterCapital(xlanguage.Chinese, "奥拉涅斯塔德")
}
