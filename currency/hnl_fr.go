//go:build (lang_fr || lang_all) && (country_all || country_americas || country_central_america || country_hn || currency_all || currency_hnl)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Hnl.RegisterName(xlanguage.French, "Lempira")
}
