//go:build country_africa || country_all || country_ls || country_southern_africa || currency_all || currency_lsl

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lsl.RegisterName(xlanguage.English, "Loti")
}
