//go:build country_all || country_asia || country_bt || country_southern_asia || currency_all || currency_btn

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BTN.RegisterName(xlanguage.English, "Ngultrum")
}
