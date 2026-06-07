//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_gi || country_southern_europe || currency_all || currency_gip)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	GIP.RegisterName(xlanguage.MustParse("zh-Hant"), "直布羅陀鎊")
}
