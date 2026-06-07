//go:build country_all || country_asia || country_south_eastern_asia || country_tl

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTimorLeste.RegisterName(xlanguage.Chinese, "东帝汶")
	dataTimorLeste.RegisterOfficialName(xlanguage.Chinese, "东帝汶民主共和国")
	dataTimorLeste.RegisterCapital(xlanguage.Chinese, "帝力")
}
