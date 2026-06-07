//go:build (lang_ar || lang_all) && (country_all || country_americas || country_pe || country_south_america || currency_all || currency_pen)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	PEN.RegisterName(xlanguage.Arabic, "سول بيروفي")
}
