//go:build country_africa || country_all || country_eastern_africa || country_zm

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZambia.RegisterName(xlanguage.Chinese, "赞比亚")
	dataZambia.RegisterOfficialName(xlanguage.Chinese, "赞比亚共和国")
	dataZambia.RegisterCapital(xlanguage.Chinese, "卢萨卡")
}
