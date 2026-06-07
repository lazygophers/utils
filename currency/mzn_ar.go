//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_mz || currency_all || currency_mzn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MZN.RegisterName(xlanguage.Arabic, "متكال موزمبيقي")
}
