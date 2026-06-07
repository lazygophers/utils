//go:build (lang_fr || lang_all) && (country_all || country_eastern_europe || country_europe || country_md || currency_all || currency_mdl)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mdl.RegisterName(xlanguage.French, "Leu moldave")
}
