//go:build country_africa || country_all || country_gm || country_western_africa || currency_all || currency_gmd

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Gmd.RegisterName(xlanguage.English, "Dalasi")
}
