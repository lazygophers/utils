//go:build country_all || country_americas || country_south_america || country_uy || currency_all || currency_uyu

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Uyu.RegisterName(xlanguage.English, "Peso Uruguayo")
}
