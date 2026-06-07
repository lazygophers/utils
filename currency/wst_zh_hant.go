//go:build (lang_zh_hant || lang_all) && (country_all || country_oceania || country_polynesia || country_ws || currency_all || currency_wst)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Wst.RegisterName(xlanguage.MustParse("zh-Hant"), "塔拉")
}
