//go:build country_ai || country_all || country_americas || country_caribbean

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAnguilla.RegisterName(xlanguage.Chinese, "安圭拉")
	dataAnguilla.RegisterOfficialName(xlanguage.Chinese, "安圭拉")
	dataAnguilla.RegisterCapital(xlanguage.Chinese, "瓦利")
}
