//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_rw || currency_all || currency_rwf)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	RWF.RegisterName(xlanguage.Spanish, "Franco ruandés")
}
