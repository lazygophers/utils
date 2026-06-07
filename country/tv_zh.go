//go:build country_all || country_oceania || country_polynesia || country_tv

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTuvalu.RegisterName(xlanguage.Chinese, "图瓦卢")
	dataTuvalu.RegisterOfficialName(xlanguage.Chinese, "图瓦卢")
	dataTuvalu.RegisterCapital(xlanguage.Chinese, "富纳富提")
}
