//go:build (lang_fr || lang_all) && (country_all || country_asia || country_sa || country_western_asia || currency_all || currency_sar)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sar.RegisterName(xlanguage.French, "Riyal saoudien")
}
