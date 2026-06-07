//go:build (lang_fr || lang_all) && (country_africa || country_all || country_middle_africa || country_st || currency_all || currency_stn)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	STN.RegisterName(xlanguage.French, "Dobra")
}
