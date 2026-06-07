//go:build (lang_ru || lang_all) && (country_all || country_europe || country_mk || country_southern_europe || currency_all || currency_mkd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mkd.RegisterName(xlanguage.Russian, "Македонский денар")
}
