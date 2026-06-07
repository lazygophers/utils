//go:build country_africa || country_all || country_sh || country_western_africa || currency_all || currency_shp

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Shp.RegisterName(xlanguage.English, "Saint Helena Pound")
}
