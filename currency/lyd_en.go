//go:build country_africa || country_all || country_ly || country_northern_africa || currency_all || currency_lyd

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	LYD.RegisterName(xlanguage.English, "Libyan Dinar")
}
