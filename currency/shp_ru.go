//go:build (lang_ru || lang_all) && (country_africa || country_all || country_sh || country_western_africa || currency_all || currency_shp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	SHP.RegisterName(xlanguage.Russian, "Фунт Святой Елены")
}
