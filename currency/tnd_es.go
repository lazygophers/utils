//go:build (lang_es || lang_all) && (country_africa || country_all || country_northern_africa || country_tn || currency_all || currency_tnd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Tnd.RegisterName(xlanguage.Spanish, "Dinar tunecino")
}
