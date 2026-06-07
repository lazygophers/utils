//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_tt || currency_all || currency_ttd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	TTD.RegisterName(xlanguage.Russian, "Доллар Тринидада и Тобаго")
}
