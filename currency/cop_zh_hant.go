//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_co || country_south_america || currency_all || currency_cop)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Cop.RegisterName(xlanguage.MustParse("zh-Hant"), "哥倫比亞披索")
}
