//go:build (lang_ar || lang_all) && (country_all || country_eastern_europe || country_europe || country_ua || currency_all || currency_uah)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	UAH.RegisterName(xlanguage.Arabic, "هريفنا أوكرانية")
}
