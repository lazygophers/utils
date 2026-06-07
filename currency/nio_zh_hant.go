//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_central_america || country_ni || currency_all || currency_nio)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Nio.RegisterName(xlanguage.MustParse("zh-Hant"), "尼加拉瓜科多巴")
}
