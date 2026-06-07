//go:build country_all || country_europe || country_gi || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGibraltar.RegisterName(xlanguage.Chinese, "直布罗陀")
	dataGibraltar.RegisterOfficialName(xlanguage.Chinese, "直布罗陀")
	dataGibraltar.RegisterCapital(xlanguage.Chinese, "直布罗陀")
}
