//go:build (lang_ar || lang_all) && (country_africa || country_all || country_ng || country_western_africa || currency_all || currency_ngn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	NGN.RegisterName(xlanguage.Arabic, "نيرة نيجيرية")
}
