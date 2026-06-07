//go:build country_all || country_europe || country_northern_europe || country_se || currency_all || currency_sek

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sek.RegisterName(xlanguage.English, "Swedish Krona")
}
