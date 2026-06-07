//go:build country_all || country_americas || country_caribbean || country_do || currency_all || currency_dop

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Dop.RegisterName(xlanguage.English, "Dominican Peso")
}
