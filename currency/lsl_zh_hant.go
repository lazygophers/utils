//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_ls || country_southern_africa || currency_all || currency_lsl)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	LSL.RegisterName(xlanguage.MustParse("zh-Hant"), "洛蒂")
}
