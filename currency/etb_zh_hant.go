//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_et || currency_all || currency_etb)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Etb.RegisterName(xlanguage.MustParse("zh-Hant"), "衣索比亞比爾")
}
