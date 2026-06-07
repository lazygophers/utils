//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_southern_africa || country_sz || currency_all || currency_szl)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SZL.RegisterName(xlanguage.MustParse("zh-Hant"), "埃馬蘭吉尼")
}
