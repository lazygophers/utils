//go:build (lang_ru || lang_all) && (country_all || country_americas || country_br || country_south_america || currency_all || currency_brl)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Brl.RegisterName(xlanguage.Russian, "Бразильский реал")
}
