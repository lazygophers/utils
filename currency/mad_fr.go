//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eh || country_ma || country_northern_africa || currency_all || currency_mad)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mad.RegisterName(xlanguage.French, "Dirham marocain")
}
