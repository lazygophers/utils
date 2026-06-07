//go:build (lang_ru || lang_all) && (country_all || country_ch || country_europe || country_li || country_western_europe || currency_all || currency_chf)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Chf.RegisterName(xlanguage.Russian, "Швейцарский франк")
}
