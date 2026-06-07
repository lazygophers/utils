//go:build country_africa || country_all || country_eastern_africa || country_mz || currency_all || currency_mzn

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mzn.RegisterName(xlanguage.English, "Mozambique Metical")
}
