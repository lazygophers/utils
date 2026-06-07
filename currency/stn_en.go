//go:build country_africa || country_all || country_middle_africa || country_st || currency_all || currency_stn

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Stn.RegisterName(xlanguage.English, "Dobra")
}
