//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_mz || currency_all || currency_mzn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mzn.RegisterName(xlanguage.MustParse("zh-Hant"), "莫三比克梅蒂卡爾")
}
