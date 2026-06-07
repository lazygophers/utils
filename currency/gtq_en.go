//go:build country_all || country_americas || country_central_america || country_gt || currency_all || currency_gtq

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	GTQ.RegisterName(xlanguage.English, "Quetzal")
}
