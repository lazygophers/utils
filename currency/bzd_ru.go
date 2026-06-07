//go:build (lang_ru || lang_all) && (country_all || country_americas || country_bz || country_central_america || currency_all || currency_bzd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BZD.RegisterName(xlanguage.Russian, "Белизский доллар")
}
