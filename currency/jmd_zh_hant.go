//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_caribbean || country_jm || currency_all || currency_jmd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Jmd.RegisterName(xlanguage.MustParse("zh-Hant"), "牙買加元")
}
