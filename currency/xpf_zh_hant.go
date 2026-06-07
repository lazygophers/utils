//go:build (lang_zh_hant || lang_all) && (country_all || country_melanesia || country_nc || country_oceania || country_pf || country_polynesia || country_wf || currency_all || currency_xpf)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Xpf.RegisterName(xlanguage.MustParse("zh-Hant"), "太平洋法郎")
}
