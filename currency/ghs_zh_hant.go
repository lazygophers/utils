//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_gh || country_western_africa || currency_all || currency_ghs)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ghs.RegisterName(xlanguage.MustParse("zh-Hant"), "迦納塞地")
}
