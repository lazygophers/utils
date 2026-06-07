//go:build country_all || country_americas || country_ca || country_northern_america || currency_all || currency_cad

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	CAD.RegisterName(xlanguage.English, "Canadian Dollar")
}
