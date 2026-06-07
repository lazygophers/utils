//go:build country_all || country_asia || country_my || country_south_eastern_asia || currency_all || currency_myr

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Myr.RegisterName(xlanguage.Chinese, "马来西亚林吉特")
}
