//go:build (lang_fr || lang_all) && (country_all || country_asia || country_kw || country_western_asia || currency_all || currency_kwd)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	KWD.RegisterName(xlanguage.French, "Dinar koweïtien")
}
