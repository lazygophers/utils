//go:build country_all || country_cy || country_europe || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCyprus.RegisterName(xlanguage.Chinese, "塞浦路斯")
	dataCyprus.RegisterOfficialName(xlanguage.Chinese, "塞浦路斯共和国")
	dataCyprus.RegisterCapital(xlanguage.Chinese, "尼科西亚")
}
