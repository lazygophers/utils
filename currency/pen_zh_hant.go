//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_pe || country_south_america || currency_all || currency_pen)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Pen.RegisterName(xlanguage.MustParse("zh-Hant"), "秘魯太陽幣")
}
