//go:build country_all || country_americas || country_south_america || country_uy

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUruguay.RegisterName(xlanguage.Chinese, "乌拉圭")
	dataUruguay.RegisterOfficialName(xlanguage.Chinese, "乌拉圭东岸共和国")
	dataUruguay.RegisterCapital(xlanguage.Chinese, "蒙得维的亚")
}
