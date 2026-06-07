//go:build (lang_es || lang_all) && (country_all || country_asia || country_lb || country_western_asia || currency_all || currency_lbp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	LBP.RegisterName(xlanguage.Spanish, "Libra libanesa")
}
