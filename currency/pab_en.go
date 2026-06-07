//go:build country_all || country_americas || country_central_america || country_pa || currency_all || currency_pab

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	PAB.RegisterName(xlanguage.English, "Balboa")
}
