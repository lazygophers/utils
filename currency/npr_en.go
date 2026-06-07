//go:build country_all || country_asia || country_np || country_southern_asia || currency_all || currency_npr

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	NPR.RegisterName(xlanguage.English, "Nepalese Rupee")
}
