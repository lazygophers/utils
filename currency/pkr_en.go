//go:build country_all || country_asia || country_pk || country_southern_asia || currency_all || currency_pkr

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	PKR.RegisterName(xlanguage.English, "Pakistan Rupee")
}
