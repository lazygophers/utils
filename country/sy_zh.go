//go:build country_all || country_asia || country_sy || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSyria.RegisterName(xlanguage.Chinese, "叙利亚")
	dataSyria.RegisterOfficialName(xlanguage.Chinese, "阿拉伯叙利亚共和国")
	dataSyria.RegisterCapital(xlanguage.Chinese, "大马士革")
}
