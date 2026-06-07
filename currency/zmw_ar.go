//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_zm || currency_all || currency_zmw)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	ZMW.RegisterName(xlanguage.Arabic, "كواشا زامبية")
}
