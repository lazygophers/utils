//go:build country_all || country_americas || country_caribbean || country_sx

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSintMaarten.RegisterName(xlanguage.Chinese, "荷属圣马丁")
	dataSintMaarten.RegisterOfficialName(xlanguage.Chinese, "荷属圣马丁")
	dataSintMaarten.RegisterCapital(xlanguage.Chinese, "菲利普斯堡")
}
