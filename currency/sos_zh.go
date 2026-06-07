//go:build country_africa || country_all || country_eastern_africa || country_so || currency_all || currency_sos

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SOS.RegisterName(xlanguage.Chinese, "索马里先令")
}
