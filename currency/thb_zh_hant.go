//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_south_eastern_asia || country_th || currency_all || currency_thb)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Thb.RegisterName(xlanguage.MustParse("zh-Hant"), "泰銖")
}
