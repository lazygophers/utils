//go:build country_all || country_melanesia || country_oceania || country_pg || currency_all || currency_pgk

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	PGK.RegisterName(xlanguage.English, "Kina")
}
