//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAruba.RegisterName(xlanguage.MustParse("zh-Hant"), "阿魯巴")
	dataAruba.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "阿魯巴")
	dataAruba.RegisterCapital(xlanguage.MustParse("zh-Hant"), "奧拉涅斯塔德")
}
