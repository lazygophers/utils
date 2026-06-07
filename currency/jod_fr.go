//go:build (lang_fr || lang_all) && (country_all || country_asia || country_jo || country_western_asia || currency_all || currency_jod)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	JOD.RegisterName(xlanguage.French, "Dinar jordanien")
}
