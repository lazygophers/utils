//go:build country_africa || country_all || country_eastern_africa || country_rw || currency_all || currency_rwf

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Rwf.RegisterName(xlanguage.English, "Rwanda Franc")
}
