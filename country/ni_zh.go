//go:build country_all || country_americas || country_central_america || country_ni

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNicaragua.RegisterName(xlanguage.Chinese, "尼加拉瓜")
	dataNicaragua.RegisterOfficialName(xlanguage.Chinese, "尼加拉瓜共和国")
	dataNicaragua.RegisterCapital(xlanguage.Chinese, "马那瓜")
}
