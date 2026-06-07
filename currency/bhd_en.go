//go:build country_all || country_asia || country_bh || country_western_asia || currency_all || currency_bhd

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BHD.RegisterName(xlanguage.English, "Bahraini Dinar")
}
