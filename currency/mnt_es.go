//go:build (lang_es || lang_all) && (country_all || country_asia || country_eastern_asia || country_mn || currency_all || currency_mnt)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mnt.RegisterName(xlanguage.Spanish, "Tugrik")
}
