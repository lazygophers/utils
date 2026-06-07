//go:build country_all || country_americas || country_caribbean || country_jm

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJamaica.RegisterName(xlanguage.Chinese, "牙买加")
	dataJamaica.RegisterOfficialName(xlanguage.Chinese, "牙买加")
	dataJamaica.RegisterCapital(xlanguage.Chinese, "金斯敦")
}
