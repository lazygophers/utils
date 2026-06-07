//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_bn || country_south_eastern_asia || currency_all || currency_bnd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BND.RegisterName(xlanguage.MustParse("zh-Hant"), "汶萊元")
}
