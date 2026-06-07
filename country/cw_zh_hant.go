//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_caribbean || country_cw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuracao.RegisterName(xlanguage.MustParse("zh-Hant"), "古拉索")
	dataCuracao.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "古拉索國")
	dataCuracao.RegisterCapital(xlanguage.MustParse("zh-Hant"), "威廉斯塔德")
}
