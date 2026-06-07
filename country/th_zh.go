//go:build country_all || country_asia || country_south_eastern_asia || country_th

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataThailand.RegisterName(xlanguage.Chinese, "泰国")
	dataThailand.RegisterOfficialName(xlanguage.Chinese, "泰王国")
	dataThailand.RegisterCapital(xlanguage.Chinese, "曼谷")
}
