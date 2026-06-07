//go:build country_all || country_americas || country_aw || country_caribbean || currency_all || currency_awg

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	AWG.RegisterName(xlanguage.English, "Aruban Florin")
}
