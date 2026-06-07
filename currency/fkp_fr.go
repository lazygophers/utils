//go:build (lang_fr || lang_all) && (country_all || country_americas || country_fk || country_south_america || currency_all || currency_fkp)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	FKP.RegisterName(xlanguage.French, "Livre des Malouines")
}
