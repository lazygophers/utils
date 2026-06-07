//go:build country_all || country_americas || country_caribbean || country_gd

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGrenada.RegisterName(xlanguage.Chinese, "格林纳达")
	dataGrenada.RegisterOfficialName(xlanguage.Chinese, "格林纳达")
	dataGrenada.RegisterCapital(xlanguage.Chinese, "圣乔治")
}
