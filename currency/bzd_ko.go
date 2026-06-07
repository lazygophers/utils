//go:build (lang_ko || lang_all) && (country_all || country_americas || country_bz || country_central_america || currency_all || currency_bzd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BZD.RegisterName(xlanguage.Korean, "벨리즈 달러")
}
