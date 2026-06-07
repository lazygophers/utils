//go:build country_all || country_americas || country_caribbean || country_kn

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintKittsAndNevis.RegisterName(xlanguage.Chinese, "圣基茨和尼维斯")
	dataSaintKittsAndNevis.RegisterOfficialName(xlanguage.Chinese, "圣基茨和尼维斯联邦")
	dataSaintKittsAndNevis.RegisterCapital(xlanguage.Chinese, "巴斯特尔")
}
