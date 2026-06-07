//go:build (lang_ja || lang_all) && (country_africa || country_all || country_bi || country_eastern_africa || currency_all || currency_bif)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BIF.RegisterName(xlanguage.Japanese, "ブルンジ・フラン")
}
