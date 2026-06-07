//go:build country_africa || country_all || country_dz || country_northern_africa || currency_all || currency_dzd

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	DZD.RegisterName(xlanguage.English, "Algerian Dinar")
}
