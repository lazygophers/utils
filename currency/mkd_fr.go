//go:build (lang_fr || lang_all) && (country_all || country_europe || country_mk || country_southern_europe || currency_all || currency_mkd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mkd.RegisterName(xlanguage.French, "Denar")
}
