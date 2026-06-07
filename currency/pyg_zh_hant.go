//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_py || country_south_america || currency_all || currency_pyg)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Pyg.RegisterName(xlanguage.MustParse("zh-Hant"), "瓜拉尼")
}
