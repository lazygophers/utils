//go:build (lang_ar || lang_all) && (country_africa || country_all || country_na || country_southern_africa || currency_all || currency_nad)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	NAD.RegisterName(xlanguage.Arabic, "دولار ناميبي")
}
