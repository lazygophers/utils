//go:build (lang_fr || lang_all) && (country_all || country_americas || country_bs || country_caribbean || currency_all || currency_bsd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	BSD.RegisterName(xlanguage.French, "Dollar bahaméen")
}
