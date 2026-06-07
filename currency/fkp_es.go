//go:build (lang_es || lang_all) && (country_all || country_americas || country_fk || country_south_america || currency_all || currency_fkp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Fkp.RegisterName(xlanguage.Spanish, "Libra malvinense")
}
