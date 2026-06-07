//go:build country_all || country_americas || country_caribbean || country_mf

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintMartin.RegisterName(xlanguage.Chinese, "法属圣马丁")
	dataSaintMartin.RegisterOfficialName(xlanguage.Chinese, "圣马丁岛集体")
	dataSaintMartin.RegisterCapital(xlanguage.Chinese, "马里戈特")
}
