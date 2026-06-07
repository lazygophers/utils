//go:build country_all || country_asia || country_np || country_southern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNepal.RegisterName(xlanguage.Chinese, "尼泊尔")
	dataNepal.RegisterOfficialName(xlanguage.Chinese, "尼泊尔联邦民主共和国")
	dataNepal.RegisterCapital(xlanguage.Chinese, "加德满都")
}
