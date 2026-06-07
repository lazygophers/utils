//go:build country_all || country_americas || country_caribbean || country_vc

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintVincentAndGrenadines.RegisterName(xlanguage.Chinese, "圣文森特和格林纳丁斯")
	dataSaintVincentAndGrenadines.RegisterOfficialName(xlanguage.Chinese, "圣文森特和格林纳丁斯")
	dataSaintVincentAndGrenadines.RegisterCapital(xlanguage.Chinese, "金斯敦")
}
