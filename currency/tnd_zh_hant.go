//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_northern_africa || country_tn || currency_all || currency_tnd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	TND.RegisterName(xlanguage.MustParse("zh-Hant"), "突尼西亞第納爾")
}
