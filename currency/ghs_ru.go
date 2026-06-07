//go:build (lang_ru || lang_all) && (country_africa || country_all || country_gh || country_western_africa || currency_all || currency_ghs)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ghs.RegisterName(xlanguage.Russian, "Ганский седи")
}
