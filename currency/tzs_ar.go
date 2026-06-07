//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_tz || currency_all || currency_tzs)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	TZS.RegisterName(xlanguage.Arabic, "شلن تنزاني")
}
