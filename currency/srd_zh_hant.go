//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_south_america || country_sr || currency_all || currency_srd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SRD.RegisterName(xlanguage.MustParse("zh-Hant"), "蘇利南元")
}
