//go:build country_all || country_cz || country_eastern_europe || country_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCzechia.RegisterName(xlanguage.Chinese, "捷克")
	dataCzechia.RegisterOfficialName(xlanguage.Chinese, "捷克共和国")
	dataCzechia.RegisterCapital(xlanguage.Chinese, "布拉格")
}
