//go:build country_africa || country_all || country_gh || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGhana.RegisterName(xlanguage.Chinese, "加纳")
	dataGhana.RegisterOfficialName(xlanguage.Chinese, "加纳共和国")
	dataGhana.RegisterCapital(xlanguage.Chinese, "阿克拉")
}
