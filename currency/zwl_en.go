//go:build country_africa || country_all || country_eastern_africa || country_zw || currency_all || currency_zwl

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	ZWL.RegisterName(xlanguage.English, "Zimbabwe Dollar")
}
