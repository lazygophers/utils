//go:build country_africa || country_all || country_eastern_africa || country_mu || currency_all || currency_mur

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mur.RegisterName(xlanguage.Chinese, "毛里求斯卢比")
}
