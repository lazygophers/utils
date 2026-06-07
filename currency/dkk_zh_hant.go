//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_dk || country_europe || country_fo || country_gl || country_northern_america || country_northern_europe || currency_all || currency_dkk)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Dkk.RegisterName(xlanguage.MustParse("zh-Hant"), "丹麥克朗")
}
