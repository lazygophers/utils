//go:build country_all || country_americas || country_central_america || country_hn

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHonduras.RegisterName(xlanguage.Chinese, "洪都拉斯")
	dataHonduras.RegisterOfficialName(xlanguage.Chinese, "洪都拉斯共和国")
	dataHonduras.RegisterCapital(xlanguage.Chinese, "特古西加尔巴")
}
