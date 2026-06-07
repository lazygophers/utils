//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_er || currency_all || currency_ern)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	ERN.RegisterName(xlanguage.Arabic, "ناكفا")
}
