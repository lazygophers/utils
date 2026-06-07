//go:build (lang_zh_hant || lang_all) && (country_all || country_melanesia || country_oceania || country_vu || currency_all || currency_vuv)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	VUV.RegisterName(xlanguage.MustParse("zh-Hant"), "瓦圖")
}
