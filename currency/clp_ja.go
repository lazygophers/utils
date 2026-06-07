//go:build (lang_ja || lang_all) && (country_all || country_americas || country_cl || country_south_america || currency_all || currency_clp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Clp.RegisterName(xlanguage.Japanese, "チリ・ペソ")
}
