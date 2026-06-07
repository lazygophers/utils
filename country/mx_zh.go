//go:build country_all || country_americas || country_central_america || country_mx

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMexico.RegisterName(xlanguage.Chinese, "墨西哥")
	dataMexico.RegisterOfficialName(xlanguage.Chinese, "墨西哥合众国")
	dataMexico.RegisterCapital(xlanguage.Chinese, "墨西哥城")
}
