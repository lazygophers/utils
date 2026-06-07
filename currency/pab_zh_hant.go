//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_central_america || country_pa || currency_all || currency_pab)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Pab.RegisterName(xlanguage.MustParse("zh-Hant"), "巴波亞")
}
