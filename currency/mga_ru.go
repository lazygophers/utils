//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_mg || currency_all || currency_mga)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mga.RegisterName(xlanguage.Russian, "Малагасийский ариари")
}
