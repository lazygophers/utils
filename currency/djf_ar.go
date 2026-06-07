//go:build (lang_ar || lang_all) && (country_africa || country_all || country_dj || country_eastern_africa || currency_all || currency_djf)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Djf.RegisterName(xlanguage.Arabic, "فرنك جيبوتي")
}
