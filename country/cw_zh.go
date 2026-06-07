//go:build country_all || country_americas || country_caribbean || country_cw

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuracao.RegisterName(xlanguage.Chinese, "库拉索")
	dataCuracao.RegisterOfficialName(xlanguage.Chinese, "库拉索国")
	dataCuracao.RegisterCapital(xlanguage.Chinese, "威廉斯塔德")
}
