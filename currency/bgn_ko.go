//go:build (lang_ko || lang_all) && (country_all || country_bg || country_eastern_europe || country_europe || currency_all || currency_bgn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bgn.RegisterName(xlanguage.Korean, "불가리아 레프")
}
