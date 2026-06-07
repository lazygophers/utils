//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_fk || country_south_america || currency_all || currency_fkp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	FKP.RegisterName(xlanguage.MustParse("zh-Hant"), "福克蘭群島鎊")
}
