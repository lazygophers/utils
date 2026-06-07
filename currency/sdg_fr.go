//go:build (lang_fr || lang_all) && (country_africa || country_all || country_northern_africa || country_sd || currency_all || currency_sdg)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SDG.RegisterName(xlanguage.French, "Livre soudanaise")
}
