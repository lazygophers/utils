//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_np || country_southern_asia || currency_all || currency_npr)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Npr.RegisterName(xlanguage.MustParse("zh-Hant"), "尼泊爾盧比")
}
