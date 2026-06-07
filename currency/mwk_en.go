//go:build country_africa || country_all || country_eastern_africa || country_mw || currency_all || currency_mwk

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mwk.RegisterName(xlanguage.English, "Malawi Kwacha")
}
