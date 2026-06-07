//go:build country_all || country_americas || country_bo || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBolivia.RegisterName(xlanguage.Chinese, "玻利维亚")
	dataBolivia.RegisterOfficialName(xlanguage.Chinese, "玻利维亚多民族国")
	dataBolivia.RegisterCapital(xlanguage.Chinese, "苏克雷")
}
