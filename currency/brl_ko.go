//go:build (lang_ko || lang_all) && (country_all || country_americas || country_br || country_south_america || currency_all || currency_brl)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Brl.RegisterName(xlanguage.Korean, "브라질 헤알")
}
