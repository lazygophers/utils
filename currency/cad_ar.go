//go:build (lang_ar || lang_all) && (country_all || country_americas || country_ca || country_northern_america || currency_all || currency_cad)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Cad.RegisterName(xlanguage.Arabic, "دولار كندي")
}
