//go:build (lang_zh_hant || lang_all) && (country_all || country_melanesia || country_oceania || country_sb || currency_all || currency_sbd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sbd.RegisterName(xlanguage.MustParse("zh-Hant"), "索羅門群島元")
}
