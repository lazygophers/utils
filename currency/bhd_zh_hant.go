//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_bh || country_western_asia || currency_all || currency_bhd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bhd.RegisterName(xlanguage.MustParse("zh-Hant"), "巴林第納爾")
}
