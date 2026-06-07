//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_tz || currency_all || currency_tzs)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Tzs.RegisterName(xlanguage.Russian, "Танзанийский шиллинг")
}
