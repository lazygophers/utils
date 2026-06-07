//go:build country_all || country_asia || country_sa || country_western_asia || currency_all || currency_sar

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SAR.RegisterName(xlanguage.English, "Saudi Riyal")
}
