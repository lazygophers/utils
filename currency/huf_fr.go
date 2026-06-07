//go:build (lang_fr || lang_all) && (country_all || country_eastern_europe || country_europe || country_hu || currency_all || currency_huf)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Huf.RegisterName(xlanguage.French, "Forint")
}
