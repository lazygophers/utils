//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_tz || currency_all || currency_tzs)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	TZS.RegisterName(xlanguage.MustParse("zh-Hant"), "坦尚尼亞先令")
}
