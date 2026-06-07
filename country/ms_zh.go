//go:build country_all || country_americas || country_caribbean || country_ms

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontserrat.RegisterName(xlanguage.Chinese, "蒙特塞拉特")
	dataMontserrat.RegisterOfficialName(xlanguage.Chinese, "蒙特塞拉特")
	dataMontserrat.RegisterCapital(xlanguage.Chinese, "布莱兹")
}
