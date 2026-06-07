//go:build (lang_es || lang_all) && (country_africa || country_all || country_lr || country_western_africa || currency_all || currency_lrd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	LRD.RegisterName(xlanguage.Spanish, "Dólar liberiano")
}
