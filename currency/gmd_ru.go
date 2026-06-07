//go:build (lang_ru || lang_all) && (country_africa || country_all || country_gm || country_western_africa || currency_all || currency_gmd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	GMD.RegisterName(xlanguage.Russian, "Даласи")
}
