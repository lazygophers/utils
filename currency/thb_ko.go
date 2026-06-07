//go:build (lang_ko || lang_all) && (country_all || country_asia || country_south_eastern_asia || country_th || currency_all || currency_thb)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	THB.RegisterName(xlanguage.Korean, "태국 바트")
}
