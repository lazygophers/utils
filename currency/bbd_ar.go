//go:build (lang_ar || lang_all) && (country_all || country_americas || country_bb || country_caribbean || currency_all || currency_bbd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bbd.RegisterName(xlanguage.Arabic, "دولار بربادوسي")
}
