//go:build (lang_ja || lang_all) && (country_all || country_eastern_europe || country_europe || country_hu || currency_all || currency_huf)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	HUF.RegisterName(xlanguage.Japanese, "フォリント")
}
