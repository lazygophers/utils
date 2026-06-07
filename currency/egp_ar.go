//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eg || country_northern_africa || currency_all || currency_egp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Egp.RegisterName(xlanguage.Arabic, "جنيه مصري")
}
