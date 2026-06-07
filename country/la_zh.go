//go:build country_all || country_asia || country_la || country_south_eastern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLaos.RegisterName(xlanguage.Chinese, "老挝")
	dataLaos.RegisterOfficialName(xlanguage.Chinese, "老挝人民民主共和国")
	dataLaos.RegisterCapital(xlanguage.Chinese, "万象")
}
