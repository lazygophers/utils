//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eh || country_ma || country_northern_africa || currency_all || currency_mad)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mad.RegisterName(xlanguage.MustParse("zh-Hant"), "摩洛哥迪拉姆")
}
