//go:build (lang_ko || lang_all) && (country_all || country_americas || country_central_america || country_pa || currency_all || currency_pab)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Pab.RegisterName(xlanguage.Korean, "발보아")
}
