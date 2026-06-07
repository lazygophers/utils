//go:build country_all || country_europe || country_gr || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreece.RegisterName(xlanguage.Chinese, "希腊")
	dataGreece.RegisterOfficialName(xlanguage.Chinese, "希腊共和国")
	dataGreece.RegisterCapital(xlanguage.Chinese, "雅典")
}
