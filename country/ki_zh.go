//go:build country_all || country_ki || country_micronesia || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKiribati.RegisterName(xlanguage.Chinese, "基里巴斯")
	dataKiribati.RegisterOfficialName(xlanguage.Chinese, "基里巴斯共和国")
	dataKiribati.RegisterCapital(xlanguage.Chinese, "塔拉瓦")
}
