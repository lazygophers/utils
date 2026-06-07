//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_mg || currency_all || currency_mga)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MGA.RegisterName(xlanguage.MustParse("zh-Hant"), "馬達加斯加阿里亞里")
}
